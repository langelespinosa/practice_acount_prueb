from fastapi import FastAPI, Query
from fastapi.responses import JSONResponse
from sentence_transformers import SentenceTransformer
import faiss
import numpy as np
import warnings
import mysql.connector
from mysql.connector import Error
import uvicorn

warnings.filterwarnings("ignore", category=FutureWarning)

DB_CONFIG = {
    'host': 'localhost',
    'database': 'email',
    'user': 'web',
    'password': 'password'
}

app = FastAPI(title="API Búsqueda Usuarios", version="1.0.0")

def obtener_usuarios_desde_mysql():
    try:
        connection = mysql.connector.connect(**DB_CONFIG)
        if connection.is_connected():
            cursor = connection.cursor(dictionary=True)
            query = "SELECT id, login, email, maildir, identificacion, grupo FROM users"
            
            cursor.execute(query)
            usuarios = cursor.fetchall()
            return usuarios
    except Error as e:
        print(f"❌ Error al conectar con MySQL: {e}")
        return []
    finally:
        if connection.is_connected():
            cursor.close()
            connection.close()

print("🔄 Cargando usuarios desde MySQL...")
usuarios = obtener_usuarios_desde_mysql()
if not usuarios:
    raise RuntimeError("No se encontraron usuarios en la base de datos")

corpus = []
id_map = {}

for i, usuario in enumerate(usuarios):
    text = f"""ID: {usuario['id']} Login: {usuario['login']} Email: {usuario['email']} Maildir: {usuario['maildir']} 
    Identificacion: {usuario['identificacion']} Grupo: {usuario['grupo']}"""
    
    corpus.append(text)
    id_map[i] = usuario["id"]

print("🔄 Generando embeddings...")
model = SentenceTransformer('sentence-transformers/all-mpnet-base-v2')
embeddings = model.encode(corpus, normalize_embeddings=True)

dimension = embeddings.shape[1]
index = faiss.IndexFlatIP(dimension)
index.add(np.array(embeddings, dtype=np.float32))
print("✅ Índice FAISS creado exitosamente")

def buscar(query, threshold=0.3):
    query_vec = model.encode([query], normalize_embeddings=True)
    total_usuarios = len(usuarios)
    D, I = index.search(np.array(query_vec, dtype=np.float32), total_usuarios)
    resultados = []
    for score, idx in zip(D[0], I[0]):
        if score >= threshold:
            resultados.append((id_map[idx], float(score)))
    resultados.sort(key=lambda x: x[1], reverse=True)
    return resultados

def buscar_hibrido(query, threshold=0.3):
    resultados = buscar(query, threshold)

    palabras = query.lower().split()

    for usuario in usuarios:
        match_exacto = False
        score_exacto = 1.0

        campos_busqueda = [
            usuario['login'], usuario['email'], usuario['maildir'],
            usuario['grupo'], usuario['identificacion']
        ]

        for palabra in palabras:
            for campo in campos_busqueda:
                if palabra in campo.lower():
                    match_exacto = True
                    break
            if match_exacto:
                break
            
        if match_exacto:
            user_id = usuario['id']
            if not any(result[0] == user_id for result in resultados):
                resultados.insert(0, (user_id, score_exacto))

    resultados.sort(key=lambda x: x[1], reverse=True)
    return resultados

def obtener_usuario_por_id(user_id):
    return next((u for u in usuarios if u["id"] == user_id), None)

@app.get("/buscar")
def endpoint_buscar(query: str = Query(..., description="Texto a buscar"), threshold: float = 0.45):
    try:
        resultados = buscar_hibrido(query, threshold)
        data = []
        for user_id, score in resultados:    
            usuario = obtener_usuario_por_id(user_id)
            if usuario:
                data.append({
                    "id": usuario["id"],
                    "login": usuario["login"],
                    "email": usuario["email"],
                    "maildir": usuario["maildir"],
                    "grupo": usuario["grupo"],
                    "identificacion": usuario["identificacion"],
                })
        return JSONResponse(data)
    
    except Exception as e:
        return JSONResponse(content={"error": str(e)}, status_code=500)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
    
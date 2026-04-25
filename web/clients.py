import requests

API_URL = "http://localhost:8080"


def _tratar_resposta(response):
    """
    Função auxiliar para validar o status code e tratar erros.
    Centraliza a lógica de mensagens para o usuário.
    """
    # 1. Sucesso (200 OK, 201 Created, 204 No Content)
    if 200 <= response.status_code < 300:
        if (
            response.status_code == 204 or not response.text
        ):  # Deletado com sucesso, sem corpo
            return {"status": "sucesso", "codigo": response.status_code}
        return response.json()

    # 2. Erros de Cliente (4xx)
    mensagem = "Erro desconhecido"

    try:
        # Tenta extrair a mensagem de erro que o seu GIN (Go) enviou
        err_json = response.json()
        # O Gin geralmente envia "error" ou "message"
        mensagem = (
            err_json.get("error")
            or err_json.get("message")
            or err_json.get("detail")
            or "Erro na requisição"
        )
    except Exception:
        # Se o backend não mandou um JSON (ex: erro 404 de rota inexistente)
        mensagem = (
            f"Erro {response.status_code}: Não foi possível processar a resposta."
        )

    # 3. Tratamento por Código Específico (Opcional, mas profissional)
    if response.status_code == 401:
        mensagem = "Não autorizado. Verifique suas credenciais."
    elif response.status_code == 404:
        mensagem = "Recurso não encontrado no servidor."
    elif response.status_code == 422:
        mensagem = "Dados inválidos. Verifique os campos enviados."
    elif response.status_code >= 500:
        mensagem = "O servidor de Backend (Go) está instável ou fora do ar."

    # Em vez de ValueError, vamos lançar uma Exception mais clara
    raise Exception(mensagem)


def _requisicao(metodo, endpoint, json=None, params=None):
    try:
        response = requests.request(
            method=metodo,
            url=f"{API_URL}{endpoint}",
            json=json,
            params=params,
            timeout=10,
        )
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de comexão: O Backend parece estar offline")
    except requests.exceptions.Timeout:
        raise Exception("Erro: A Api demorou demais para responder.")


def cadastrar_usuario(usuario_dados: dict):
    return _requisicao("POST", "/usuario", json=usuario_dados)


def buscar_usuario(username: str):
    return _requisicao("GET", "/usuario/buscar", params={"username": username})


def atualizar_usuario(usuario_dados: dict):
    return _requisicao("PUT", "/usuario", json=usuario_dados)


def deletar_usuario(id: int):
    return _requisicao("DELETE", f"/usuario/{id}")


def auth_usuario(usuario_login_password: dict):
    return _requisicao("POST", "/usuario/auth", json=usuario_login_password)

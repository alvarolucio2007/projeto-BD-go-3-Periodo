import streamlit as st
import pandas as pd
import warnings

warnings.filterwarnings("ignore", category=DeprecationWarning)
warnings.filterwarnings("ignore", message=".*use_container_width.*")

# Configuração da página (deve ser a primeira linha do streamlit)
st.set_page_config(page_title="Meu Primeiro Front", layout="wide")

# --- BARRA LATERAL (SIDEBAR) ---
with st.sidebar:
    st.title("Escola Go")
    opcao = st.radio("Navegação", ["Início", "Usuários", "Provas"])
    st.divider()
    if st.button("Sair", type="primary"):
        st.write("Logout efetuado")

# --- CONTEÚDO PRINCIPAL ---
if opcao == "Início":
    st.header("Bem-vindo ao Sistema")
    st.write("Isso aqui substitui todo aquele HTML complexo.")

    # Um card simples de métrica
    st.metric(label="Usuários Ativos", value="15", delta="2 novos")

elif opcao == "Usuários":
    st.header("Gerenciamento de Usuários")

    # Simulando dados que viriam do seu gRPC
    dados = [
        {"id": 1, "nome": "Alvaro", "cargo": "Admin"},
        {"id": 2, "nome": "Joao", "cargo": "Aluno"},
    ]
    df = pd.DataFrame(dados)

    # A tabela que você quase quebrou a cabeça para fazer em HTML:
    st.subheader("Lista Interativa")
    st.data_editor(df, use_container_width=True)

    # Botão para adicionar (abre um modal/expander)
    with st.expander("Cadastrar Novo Usuário"):
        nome = st.text_input("Nome do Usuário")
        cargo = st.selectbox("Cargo", ["Admin", "Professor", "Aluno"])
        if st.button("Salvar no Banco"):
            st.success(f"Usuário {nome} enviado via gRPC!")

import streamlit as st
from frontend import FrontEnd


def main():
    st.set_page_config(page_title="Sistema Escolar", page_icon="🎓", layout="wide")

    app = FrontEnd()

    # Se não houver 'logado' no estado da sessão, mostra Login
    if "logado" not in st.session_state or not st.session_state["logado"]:
        app.renderizar_login()
    else:
        # Usuário logado: mostra o sistema real
        opcao = app.renderizar_menu_lateral()

        if opcao == "Usuários":
            app.renderizar_usuarios()
        elif opcao == "Provas":
            app.renderizar_provas()
        elif opcao == "Notas":
            app.renderizar_notas()
        elif opcao == "Relatórios":
            app.renderizar_relatorios()

        # Botão de Logout no final do menu lateral
        if st.sidebar.button("Sair"):
            st.session_state["logado"] = False
            st.rerun()


if __name__ == "__main__":
    main()

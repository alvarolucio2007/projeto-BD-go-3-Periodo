import streamlit as st
import clients as ct
import pandas as pd
import plotly.express as px


class FrontEnd:
    def renderizar_login(self) -> None:
        st.title("🔐 Acesso ao Sistema")

        with st.container():
            col1, col2, col3 = st.columns([1, 2, 1])
            with col2:
                with st.form("form_login"):
                    user = st.text_input("Usuário")
                    pw = st.text_input("Senha", type="password")
                    entrar = st.form_submit_button("Entrar", use_container_width=True)

                    if entrar:
                        try:
                            # Chama o seu tradutor Python que bate no Go
                            res = ct.auth_usuario({"username": user, "password": pw})

                            # Supondo que seu AuthResult retorne {status: bool, role: string}
                            if res.get("auth"):
                                st.session_state["logado"] = True
                                st.session_state["usuario"] = user
                                st.session_state["role"] = res.get("role")
                                st.rerun()
                            if not res.get("auth"):
                                st.error(st.session_state["mensagem"])
                        except Exception as e:
                            st.error(f"Erro ao autenticar: {e}")

    def renderizar_menu_lateral(self) -> str:
        with st.sidebar:
            st.header(f"Olá, {st.session_state.get('usuario', 'Usuário')}!")
            st.divider()

            opcao = st.radio(
                "Selecione a ação",
                ("Usuários", "Provas", "Notas", "Relatórios", "Dashboards"),
            )
            return opcao

    def renderizar_usuarios(self) -> None:
        st.title("👥 Gestão de Usuários")

        # Criando as abas para o CRUD
        aba_listar, aba_cadastrar, aba_atualizar, aba_excluir = st.tabs(
            ["Listar/Buscar", "Cadastrar", "Atualizar", "Excluir"]
        )

        # --- READ (LISTAR/BUSCAR) ---
        with aba_listar:
            st.subheader("Buscar Usuário")
            col1, col2 = st.columns([3, 1])
            with col1:
                username_busca = st.text_input("Username para busca", key="search_user")
            with col2:
                btn_buscar = st.button("🔍 Buscar", use_container_width=True)

            if btn_buscar:
                try:
                    # Usando seu ct.buscar_usuario(username) -> GET /usuario/buscar
                    res = ct.buscar_usuario(username_busca)
                    # Se o seu Go retorna um objeto único ou lista, trate aqui
                    if res:
                        st.table(res)  # Ou st.dataframe(res)
                    else:
                        st.info("Nenhum usuário encontrado.")
                except Exception as e:
                    st.error(f"Erro: {e}")

        # --- CREATE (CADASTRAR) ---
        with aba_cadastrar:
            st.subheader("Novo Cadastro")
            with st.form("form_add_user"):
                new_user = st.text_input("Username")
                new_pw = st.text_input("Senha", type="password")
                new_role = st.selectbox("Papel (Role)", ["admin", "professor", "aluno"])

                if st.form_submit_button("Salvar Usuário"):
                    try:
                        ct.cadastrar_usuario(
                            {"username": new_user, "password": new_pw, "role": new_role}
                        )
                        st.success(f"Usuário {new_user} criado!")
                    except Exception as e:
                        st.error(e)

        # --- UPDATE (ATUALIZAR) ---
        with aba_atualizar:
            st.subheader("Editar Usuário")
            # Dica: No Update, você precisa do ID que o Go exige
            up_id = st.number_input("ID do Usuário para editar", min_value=1, step=1)
            with st.form("form_update_user"):
                up_user = st.text_input("Novo Username")
                up_pw = st.text_input("Nova Senha", type="password")
                up_role = st.selectbox("Novo Role", ["admin", "professor", "aluno"])

                if st.form_submit_button("Atualizar"):
                    try:
                        ct.atualizar_usuario(
                            {
                                "id": int(up_id),
                                "username": up_user,
                                "password": up_pw,
                                "role": up_role,
                            }
                        )
                        st.success("Dados atualizados com sucesso!")
                    except Exception as e:
                        st.error(e)

        # --- DELETE (EXCLUIR) ---
        with aba_excluir:
            st.subheader("Remover Usuário")
            del_id = st.number_input(
                "ID do Usuário para deletar", min_value=1, step=1, key="del_id"
            )
            if st.button("❌ EXCLUIR DEFINITIVAMENTE", type="primary"):
                try:
                    ct.deletar_usuario(del_id)
                    st.success("Usuário removido do sistema.")
                except Exception as e:
                    st.error(e)

    def renderizar_provas(self) -> None:
        st.title("📝 Gestão de Provas")

        aba_listar, aba_cadastrar, aba_atualizar, aba_excluir = st.tabs(
            ["Ver Todas", "Criar Prova", "Atualizar", "Remover"]
        )

        # --- READ (LISTAR) ---
        with aba_listar:
            if st.button("🔄 Atualizar Lista", key="btn_update_provas"):
                res = ct.ler_todas_provas()
                provas = res.get("provas", []) if isinstance(res, dict) else res

                if provas:
                    df = pd.DataFrame(provas)

                    # Mapeamento para nomes amigáveis
                    mapeamento_provas = {
                        "id": "ID",
                        "nome_prova": "Nome",
                        "turma_prova": "Turma",
                        "materia_prova": "Matéria",
                        "data_prova": "Data da Prova",
                    }

                    # Seleciona e renomeia apenas se a coluna existir no JSON
                    cols = [c for c in mapeamento_provas.keys() if c in df.columns]
                    df = df[cols].rename(columns=mapeamento_provas)

                    # Formatação de Data/Hora
                    if "Data da Prova" in df.columns:
                        df["Data da Prova"] = pd.to_datetime(
                            df["Data da Prova"]
                        ).dt.strftime("%d/%m/%Y %H:%M")

                    st.dataframe(df, use_container_width=True, hide_index=True)
                else:
                    st.info("Nenhuma prova cadastrada.")
        # --- CREATE (CADASTRAR) ---
        with aba_cadastrar:
            with st.form("form_add_prova"):
                nome_prova = st.text_input("Nome da Avaliação (ex: P1, Simulado)")
                turma_prova = st.text_input("Turma da avaliação")
                materia_prova = st.text_input("Matéria da avaliação")
                col1, col2 = st.columns(2)
                with col1:
                    data_p = st.date_input("Data")
                with col2:
                    hora_p = st.time_input("Horário")

                if st.form_submit_button("Salvar Prova"):
                    try:
                        ct.cadastrar_prova(
                            {
                                "nome_prova": nome_prova,
                                "turma_prova": turma_prova,
                                "materia_prova": materia_prova,
                                "data_prova": f"{data_p}T{hora_p}Z",  # Go espera string ou parse de data
                            }
                        )
                        st.success(f"Prova '{nome_prova}' cadastrada!")
                    except Exception as e:
                        st.error(e)

        # --- UPDATE (ATUALIZAR) ---
        with aba_atualizar:
            up_id = st.number_input(
                "ID da Prova", min_value=1, step=1, key="up_prova_id"
            )
            with st.form("form_update_prova"):
                up_nome = st.text_input("Nome da Avaliação")
                up_turma = st.text_input("Turma da avaliação")
                up_materia = st.text_input("Matéria da avaliação")
                col1, col2 = st.columns(2)
                with col1:
                    up_data = st.date_input("Data")
                with col2:
                    up_hora = st.time_input("Horário")

                if st.form_submit_button("Atualizar Dados"):
                    try:
                        ct.atualizar_prova(
                            {
                                "id": int(up_id),
                                "nome_prova": up_nome,
                                "turma_prova": up_turma,
                                "materia_prova": up_materia,
                                "data_prova": f"{up_data}T{up_hora}Z",
                            }
                        )
                        st.success("Prova atualizada!")
                    except Exception as e:
                        st.error(e)

        # --- DELETE (EXCLUIR) ---
        with aba_excluir:
            del_id = st.number_input(
                "ID da Prova para remover", min_value=1, step=1, key="del_prova_id"
            )
            if st.button("🗑️ Excluir Prova", type="primary"):
                try:
                    ct.deletar_prova(del_id)
                    st.success("Prova removida com sucesso.")
                except Exception as e:
                    st.error(e)

    def renderizar_notas(self) -> None:
        st.title("📊 Gestão de Notas")

        # Criando abas idênticas ao padrão de Provas
        aba_listar, aba_lancar, aba_atualizar, aba_excluir = st.tabs(
            ["Ver Notas", "Lançar Nota", "Atualizar", "Remover"]
        )

        # --- READ (LISTAR NOTAS) ---
        with aba_listar:
            col_busca, col_btn = st.columns([3, 1])
            with col_busca:
                aluno_busca = st.text_input(
                    "Nome do Aluno",
                    placeholder="Insira o nome do aluno",
                    key="busca_nota_nome",
                )
            with col_btn:
                st.write("##")  # Alinhamento vertical
                btn_buscar = st.button("🔍 Buscar", use_container_width=True)

            if btn_buscar:
                try:
                    res = ct.buscar_nota(aluno_busca)
                    notas = res.get("notas", []) if isinstance(res, dict) else res

                    if notas:
                        df = pd.DataFrame(notas)

                        # Mapeamento padrão para Notas
                        mapeamento_notas = {
                            "id": "ID Registro",
                            "username": "Aluno",
                            "nome_prova": "Avaliação",
                            "nota_prova": "Nota",
                            "data_aplicacao": "Data da Prova",
                        }

                        # Filtra colunas existentes e renomeia
                        cols = [c for c in mapeamento_notas.keys() if c in df.columns]
                        df = df[cols].rename(columns=mapeamento_notas)

                        # Formatação de Data/Hora
                        if "Data da Prova" in df.columns:
                            df["Data da Prova"] = pd.to_datetime(
                                df["Data da Prova"]
                            ).dt.strftime("%d/%m/%Y %H:%M")

                        # Exibição com destaque na maior nota
                        st.dataframe(
                            df.style.highlight_max(
                                axis=0, subset=["Nota"], color="#2e7b32"
                            )
                            if "Nota" in df.columns
                            else df,
                            use_container_width=True,
                            hide_index=True,
                        )
                    else:
                        st.info(f"Nenhuma nota encontrada para '{aluno_busca}'.")
                except Exception as e:
                    st.error(f"Erro ao carregar notas: {e}")  # --- CREATE (LANÇAR) ---
        with aba_lancar:
            try:
                # Busca dinâmica para os Selectboxes
                res_u = ct.buscar_usuario("")
                res_p = ct.ler_todas_provas()

                usuarios = (
                    res_u if isinstance(res_u, list) else res_u.get("usuarios", [])
                )
                provas = res_p.get("provas", []) if isinstance(res_p, dict) else res_p

                if not usuarios or not provas:
                    st.warning("É necessário ter alunos e provas cadastrados primeiro.")
                else:
                    with st.form("form_lancar_nota"):
                        aluno = st.selectbox(
                            "Selecione o Aluno",
                            options=usuarios,
                            format_func=lambda u: f"{u['username']} (ID: {u['id']})",
                        )

                        prova = st.selectbox(
                            "Selecione a Prova",
                            options=provas,
                            format_func=lambda p: (
                                f"{p['nome_prova']} - {p['materia_prova']} (ID: {p['id']})"
                            ),
                        )

                        # Number input para garantir que o Python trate como float
                        nota_valor = st.number_input(
                            "Valor da Nota",
                            min_value=0.0,
                            max_value=10.0,
                            step=0.1,
                            format="%.1f",
                        )

                        if st.form_submit_button("Confirmar Lançamento"):
                            payload = {
                                "usuario_id": int(aluno["id"]),
                                "prova_id": int(prova["id"]),
                                "nota_prova": float(nota_valor),
                            }
                            ct.cadastrar_nota(payload)
                            st.success(
                                f"Nota {nota_valor} atribuída a {aluno['username']}!"
                            )
            except Exception as e:
                st.error(f"Erro na interface de lançamento: {e}")

        # --- UPDATE (ATUALIZAR) ---
        with aba_atualizar:
            st.subheader("📝 Editar Nota")
            with st.form("form_update_nota"):
                up_nota_id = st.number_input(
                    "ID da Nota (Registro)", min_value=1, step=1
                )
                nova_nota_val = st.number_input(
                    "Novo Valor da Nota", min_value=0.0, max_value=10.0, step=0.1
                )

                if st.form_submit_button("Salvar Alteração"):
                    try:
                        # Aqui passamos o ID da nota e o novo valor
                        ct.editar_nota(
                            {"id": int(up_nota_id), "nota_prova": float(nova_nota_val)}
                        )
                        st.success(
                            f"Nota ID {up_nota_id} atualizada para {nova_nota_val}!"
                        )
                    except Exception as e:
                        st.error(f"Erro ao atualizar: {e}")

        # --- DELETE (EXCLUIR) ---
        with aba_excluir:
            st.subheader("🗑️ Remover Registro de Nota")
            del_nota_id = st.number_input(
                "ID da Nota para remover", min_value=1, step=1, key="del_nota_id_input"
            )

            if st.button("Confirmar Exclusão Definitiva", type="primary"):
                try:
                    ct.deletar_nota(int(del_nota_id))
                    st.success(f"Registro {del_nota_id} removido.")
                    # st.rerun() # Opcional: recarrega para limpar a tela
                except Exception as e:
                    st.error(f"Erro ao deletar: {e}")

    def renderizar_relatorios(self) -> None:
        st.title("📈 Relatórios e Análises")
        aba_aproveitamento, aba_panorama = st.tabs(
            ["🏆 Boletim de Aproveitamento", "👥 Panorama Geral de Alunos"]
        )
        with aba_aproveitamento:
            if st.button("Gerar Relatório de Aproveitamento"):
                dados = ct.inner_join()
                # CHAVE CORRETA: "resultado" (conforme seu JSON)
                lista = dados.get("resultado", []) if isinstance(dados, dict) else dados
                if lista:
                    df = pd.DataFrame(lista)

                    # Mapeamento para garantir a ordem e nomes bonitos
                    # As chaves aqui devem ser EXATAMENTE as do seu JSON (minúsculas)
                    mapeamento = {
                        "username": "Aluno",
                        "nome_prova": "Avaliação",
                        "nota_prova": "Nota",
                        "data_aplicacao": "Data da Prova",
                    }

                    # Reordenar e renomear
                    df = df[list(mapeamento.keys())].rename(columns=mapeamento)

                    # Convertendo a data de string para objeto datetime (mais bonito)
                    df["Data da Prova"] = pd.to_datetime(
                        df["Data da Prova"]
                    ).dt.strftime("%d/%m/%Y %H:%M")

                    st.dataframe(
                        df.style.highlight_max(
                            axis=0, subset=["Nota"], color="#2e7b32"
                        ),
                        use_container_width=True,
                        hide_index=True,
                    )

                    st.metric("Média Geral", f"{df['Nota'].mean():.2f}")
                else:
                    st.info("Nenhum dado encontrado.")

        with aba_panorama:
            st.subheader("Status de Todos os Alunos")
            st.caption("Mostra todos os alunos cadastrados e suas notas (se houver).")

            if st.button("Gerar Panorama Geral"):
                try:
                    # Chama o Left Join no conector
                    dados = ct.left_join()

                    # Extrai a lista usando a chave confirmada "resultado"
                    lista = (
                        dados.get("resultado", []) if isinstance(dados, dict) else dados
                    )

                    if lista:
                        df = pd.DataFrame(lista)

                        # Mapeamento para renomear colunas
                        mapeamento = {
                            "username": "Aluno",
                            "nome_prova": "Avaliação",
                            "nota_prova": "Nota",
                            "data_aplicacao": "Data da Prova",
                        }

                        # Garante que as colunas existam antes de renomear
                        colunas_existentes = [
                            c for c in mapeamento.keys() if c in df.columns
                        ]
                        df = df[colunas_existentes].rename(columns=mapeamento)

                        # --- TRATAMENTO DE VALORES NULOS (O diferencial do Left Join) ---
                        # Onde estiver NaN (null vindo do Go), colocamos avisos claros
                        df["Avaliação"] = df["Avaliação"].fillna("🚫 Não Realizada")
                        df["Nota"] = df["Nota"].fillna("🚫 Não Realizada")

                        # Formata a data apenas se ela não for nula
                        df["Data da Prova"] = (
                            pd.to_datetime(df["Data da Prova"])
                            .dt.strftime("%d/%m/%Y %H:%M")
                            .fillna("---")
                        )

                        # Exibição com Estilo
                        st.table(df)  # st.table é ótimo para panoramas estáticos

                        # Insight rápido
                        total_alunos = len(df["Aluno"].unique())
                        alunos_sem_nota = df[
                            df["Avaliação"] == "🚫 Não Realizada"
                        ].shape[0]

                        c1, c2 = st.columns(2)
                        c1.metric("Total de Alunos", total_alunos)
                        c2.metric(
                            "Pendências de Prova",
                            alunos_sem_nota,
                            delta_color="inverse",
                        )
                    else:
                        st.warning("Nenhum aluno encontrado na base de dados.")
                except Exception as e:
                    st.error(f"Erro ao gerar panorama: {e}")

    def renderizar_dashboards(self) -> None:
        st.title("📊 Dashboards de Desempenho")

        # Criando as abas
        aba_barra, aba_dispersao, aba_horizontal, aba_pizza = st.tabs(
            ["Gráfico de Barra", "Dispersão", "Barra Horizontal", "Pizza"]
        )

        # --- ABA 1: BARRA ---
        with aba_barra:
            st.subheader("Análise Quantitativa")
            texto = st.text_input(
                "Procurar aluno para gerar o gráfico (Deixe em branco para resultado geral)"
            )
            dados = ct.dashboard_quantidade_prova(texto)
            df = pd.DataFrame(dados.items(), columns={"Quantidade"})
            fig_barra = px.bar(
                df,
                x="Item",
                y="Quantidade",
                color="Item",
                title="Quantidade por Categoria/Prova",
                template="plotly_white",
            )
            st.plotly_chart(fig_barra, use_container_width=True)

        # --- ABA 2: DISPERSÃO ---
        with aba_dispersao:
            st.subheader("Análise de Dispersão")
            texto = st.text_input(
                "Procurar aluno para gerar o gráfico (Deixe em branco para resultado geral)"
            )
            dados = ct.dashboard_quantidade_nota_prova(texto)
            df = pd.DataFrame(dados.items(), columns={"Quantidade", "Média"})
            fig_dispersao = px.scatter(
                df,
                x="Item",
                y="Quantidade",
                size="Quantidade",
                color="Item",
                title="Dispersão de Quantidade",
                template="plotly_white",
            )
            st.plotly_chart(fig_dispersao, use_container_width=True)

        # --- ABA 3: BARRA HORIZONTAL ---
        with aba_horizontal:
            st.subheader("Ranking / Comparativo")
            texto = st.text_input(
                "Procurar aluno para gerar o gráfico (Deixe em branco para resultado geral)"
            )
            dados = ct.dashboard_media_nota_materia(texto)
            df = pd.DataFrame(dados.items(), columns={"Quantidade", "Média"})
            fig_horizontal = px.bar(
                df,
                x="Quantidade",
                y="Item",
                orientation="h",  # Define o gráfico como horizontal
                color="Item",
                title="Comparativo Horizontal",
                template="plotly_white",
            )
            # Dica: Inverter o eixo Y para o maior/primeiro item ficar no topo
            fig_horizontal.update_yaxes(autorange="reversed")
            st.plotly_chart(fig_horizontal, use_container_width=True)

        # --- ABA 4: PIZZA ---
        with aba_pizza:
            st.subheader("Distribuição Percentual")
            texto = st.text_input(
                "Procurar aluno para gerar o gráfico (Deixe em branco para resultado geral)"
            )
            dados = ct.dashboard_distribuicao_status()
            df = pd.DataFrame(dados.items(), columns={"Resultado"})
            fig_pizza = px.pie(
                df,
                names="Item",
                values="Quantidade",
                title="Proporção dos Dados",
                hole=0.3,  # Cria um gráfico de rosca (opcional, mas fica mais moderno)
            )
            st.plotly_chart(fig_pizza, use_container_width=True)

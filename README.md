# Progetto provvisorio Antares

# 1) L'applicazione è avviabile su un server gestito dal framework echo
# 2) L'applicazione risponde ad una chiamata API che legge un csv in entrata
# 3) L'applicazione (se il device rivelato risulta essere MODBUS) trasmette il messaggio ad un server TCP
# 4) Il server TCP risponde visualizzando il messaggio ricevuto in precedenza

# *Da risolvere

# 1) L'applicazione gestisce le connessioni tcp senza che provi ogni volta a generarne una nuova (utilizza la connessione già aperta)
# 2) Il comando di READ va lanciato senza l'interazione con il terminale

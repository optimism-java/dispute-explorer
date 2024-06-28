FROM postgres:15

RUN apt-get update &&\
    apt-get -y install make && \
    apt-get -y install gcc && \
    apt-get -y install postgresql-15-wal2json && \
    apt-get -y install postgresql-server-dev-15 wget && \
    wget https://github.com/eulerto/wal2json/archive/refs/tags/wal2json_2_6.tar.gz &&\
    tar -zxf wal2json_2_6.tar.gz && \
    cd wal2json-wal2json_2_6 && \
    export PATH=/usr/lib/postgresql/14/bin:$PATH && \
    make && make install

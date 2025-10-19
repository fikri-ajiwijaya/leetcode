FROM debian:sid

RUN apt-get update && \
    apt-get -y install locales && \
    sed -i '/en_US.UTF-8/s/^# //g' '/etc/locale.gen' && \
    locale-gen

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN apt-get -y dist-upgrade && \
    apt-get -y autopurge && \
    apt-get -y install curl git gnupg less man patch ssh vim

RUN apt-get -y install clang cmake make

RUN architecture="$(uname -m)" && \
    case $architecture in \
    	x86_64) architecture="amd64";; \
    	aarch64 | armv8*) architecture="arm64";; \
    	aarch32 | armv7* | armvhf*) architecture="armv6l";; \
    	i?86) architecture="386";; \
    	*) echo "(!) Architecture $architecture unsupported"; exit 1;; \
    esac && \
    version="$(curl -s 'https://go.dev/VERSION?m=text' | head -n1)" && \
    curl -L "https://go.dev/dl/${version}.linux-${architecture}.tar.gz" | tar -xzC '/usr/local' && \
    echo 'export PATH="$PATH:/usr/local/go/bin"' > '/etc/profile.d/golang.sh'

COPY --chmod=644 '.vimrc' '/etc/skel'
ARG username='dev'
RUN sed -i '/PS1.*\\\$/s/\\\$/\\n\0/' '/etc/skel/.bashrc' && \
    useradd "${username}" -ms '/bin/bash'

ARG workspaceFolder
RUN install -d -o "${username}" -g "$(id -gn "${username}")" "${workspaceFolder}" && \
    ln -s "${workspaceFolder}/host/.vscode" "${workspaceFolder}/.vscode"

USER "${username}"

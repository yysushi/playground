FROM python:3.8.6-buster
RUN pip install google-auth \
                google-cloud-container google-cloud-compute \
                kubernetes
ENV CLOUDSDK_CORE_DISABLE_PROMPTS=1
ENV CLOUDSDK_INSTALL_DIR=/
RUN curl https://sdk.cloud.google.com | bash
ENV PATH $PATH:/google-cloud-sdk/bin
WORKDIR /google-cloud-container
COPY . .

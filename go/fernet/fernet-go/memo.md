# Fernet

## Cli by Python

```
virtualenv venv
source venv/bin/activate
pip install cryptography
python -c 'import sys, cryptography.fernet; print(cryptography.fernet.Fernet(sys.argv[1]).encrypt(b"A really secret message. Not for prying eyes."))' 'QE2EoUQm7y0lRWL4-pfPgTSK-4S0bCwEXSb3NKeQbgc='
```

## Cli by Golang

```
go install github.com/fernet/fernet-go/cmd/fernet-keygen@latest
go install github.com/fernet/fernet-go/cmd/fernet-sign@latest
```

```
export KEY=$(fernet-keygen)
echo "text" fernet-sign KEY
```

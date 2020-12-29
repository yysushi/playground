import ldap


def main():
    s = ldap.initialize("ldaps://ldap-server.example.org")
    # s.set_option(ldap.OPT_X_TLS_CACERTDIR, '/path/to/ca.pem')
    print(s.get_option(ldap.OPT_X_TLS_CACERTDIR))
    s.start_tls_s()


if __name__ == '__main__':
    main()

def wei():
    all_raw_servers = iter(([{'name': 'hoge'}], [{'name': 'fuga', 'id': 'me'}]))
    server_name = 'fuga'
    project_names = iter(('zemi', 'zemi2'))

    def server_id_filter(raw_servers):
        return filter(lambda x: x['name'] == server_name, raw_servers)
    all_raw_servers = map(server_id_filter, all_raw_servers)
    all_raw_servers = zip(all_raw_servers, project_names)
    for raw_servers, project_name in all_raw_servers: 
        try:
            raw_server = next(raw_servers)
        except StopIteration:
            continue
        return project_name, raw_server['id']
    raise Exception("hoge")


print(wei())

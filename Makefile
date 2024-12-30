ssh:
	#@ssh-keygen -f "/Users/uki/.ssh/known_hosts" -R "[localhost]:2222"
	@cat ssh_http_tunel.go | ssh localhost -p 2222
	
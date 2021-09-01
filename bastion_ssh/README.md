# bastion_ssh

This is a long command that will create a bastion session in OCI Bastions Service, generate the ssh command, and run it in your terminal. Works on Mac, probably would work on most Linux distros, and may work got git-bash. I haven't done much testing beyond what has been beneficial for my own use so far, so fair warning.

Remember to set the variables with the OCID of your bastion, OCID of your compute, and public/private key pair locations. If you want to log in as a user besides _opc_ or have a session shorter/longer than 30 minutes just edit _--target-os-username_ and _--session-ttl_ as needed.

Managed SSH sessions only right now, so make sure that the bastion agent is running prior to trying to connect.

Good luck!

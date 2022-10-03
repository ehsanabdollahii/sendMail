Let’s break down the above command and its directives:

* docker run starts a docker container.
* -d directive runs the docker container in the background.
* -p 1025:1025 and -p 8025:8025 directives expose MailHog’s default SMTP ports to your local ones.
* mailhog/mailhog grabs the latest MailHog build from Docker Hub
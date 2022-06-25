# Rest
Rest in Go

Es un proyecto que se basa en desarrollar un CRUD en go

Para probarlo debes iniciarlo en docker 

Sobre la carpeta database
# "docker build . -t 'database'"
luego
# "docker run -p 54321:5432 'database'"

Seguimos 

Sobre la carpeta principal
# "docker build . -t 'Nombre'"
luego
# "docker run -p 5050:5050 'Nombre'"

y eso es todo, podras probarlo desde postman 

(Todas las rutas estan en el archivo main.go)
El proyecto consta de usuarios y post, puedes crear usuarios, iniciar sesion, postear, leer post, actualizarlos, incluso eliminarlos.
Con la ayuda de los websockets, mientras estes conectado, si otra persona postea, te llegara de inmediato.

Cualquier inquietud no dude en contactarse conmigo en kgpicon@hotmail.com

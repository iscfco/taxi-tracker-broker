Publish-Subscribe Broker V1

Caracteristicas:
* Basado en websockets
* Soporta el patron de mensajes Publish-Subscribe
* Basado en tópicos
* El intercambio de mensaje es sobre el formato JSON



Limitantes:
* Actualmente no se ejecuta algun tipo de autenticacion, se pretende que la autenticacion sea basa en tokens, este feature esta por desarrollarse
* Los clientes se encargan de persistir las conexiones, es decir si en una webpage se recarga la página la conexión se cerrará y los clientes ejecutaran la tarea de suscribirse nuevamente al topico, se requieren de librerias clientes para ejecutar ese tipo de persistencia de manera abstracta.
* El nivel de QoS por el momento es solamente 0

# CulturaGRX

## Problema

Este problema lo tienen muchos ciudadanos de la ciudad de Granada, al igual que muchos turistas que vienen a visitar la ciudad y provincia. Cuando quieren, en su tiempo libre, vacaciones o fines de semana, o en el caso de los turistas en su visita, no pueden ni saben descubrir ni visitar muchísimos espacios culturales y patrimoniales, ya que la información de la que disponemos de su existencia es prácticamente nula. Además, tampoco saben sus horarios de visita, la historia que tienen detrás esos espacios, etc.

Por otro lado, cuando se tienen pocos días y unos intereses concretos, como el arte islámico, no saben qué sitios pueden ver ni cuándo, porque cada uno tiene un horario distinto que además cambia según la temporada. Por eso acaban encontrándose sitios cerrados o perdiéndose los que más les interesan.

Es un problema porque muchos de estos espacios reciben pocas visitas, ya que muchas personas no saben ni de su existencia.

## Fuentes de datos

El problema necesita información sobre los diferentes bienes culturales y patrimoniales de Granada y provincia, que son joyas que están sin descubrir y por falta de información no conocemos ni están al alcance de nosotros. De diferentes fuentes abiertas podremos obtener información tanto de sus ubicaciones, historia, etc.

La primera que voy a usar es esta fuente de datos para el proyecto, el primero es de la web de la Junta de Andalucía y su Instituto Andaluz de Patrimonio Histórico, esta fuente la voy a utilizar para la parte de qué espacios existen y dónde están,  en el cual dentro de esta web: [portal de datos abiertos de la Junta](https://www.juntadeandalucia.es/datosabiertos/portal/dataset/patrimonio-inmueble-de-andalucia) nos encontramos con una API con formato XML, CSV y JSON para aportar id, codigo, denominacion, provincia, municipio y caracterizacion. En un ejemplo de la API (json): [API en formato JSON](https://www.juntadeandalucia.es/datosabiertos/portal/iaph/dataset/bien/inmueble?page=0&rows=31000&format=json)

Aquí podemos extraer los datos que he comentado antes, su identificación, su código (que como podemos observar todos los de Granada empiezan por 0118), denominación, el municipio y la caracterización.

Además de por cada bien patrimonial, tenemos dentro de su página personalizada por ejemplo de la iglesia de la Encarnación de Albolote, con la siguiente URL:  [ficha de la Iglesia de la Encarnación](https://guiadigital.iaph.es/bien/inmueble/9232/granada/albolote/iglesia-de-la-encarnacion) dentro de la misma podemos descargar arriba a la derecha los datos abiertos de cada uno en formato .jsonld.:


Dentro de cada ficha podemos extraer las coordenadas, la protección, el periodo y la descripción.

Para poder acceder a estos datos abiertos, sería necesario entrar en la URL específica de cada bien patrimonial, que se ubica un botón justo en la esquina derecha llamado Descargar datos abiertos, esto lo he conseguido buscando en las herramientas del desarrollador, dentro de Network y escribiendo dentro de Filter el id del bien que quieras encontrar y pulsando en Headers encuentras la URL pertinente, por ejemplo: `https://guiadigital.iaph.es/api/1.0/bien/inmueble/enriquecido/9232` , como vemos el último número (id), es el que hacer referencia a uno concreto, pero al abrirla directamente da error,  porque la API exige un token de acceso. La web de la ficha llama a esa URL enviando un token; por eso en las herramientas de desarrollador la petición funciona, pero si abres la URL tú solo, sin token, da error. Apareciendo el siguiente mensaje: 

```
<ams:fault xmlns:ams="http://wso2.org/apimanager/security">
<ams:code>900902</ams:code>
<ams:message>Missing Credentials</ams:message>
<ams:description>Required OAuth credentials not provided. Make sure your API invocation call has a header: "Authorization: Bearer ACCESS_TOKEN"</ams:description>
</ams:fault>
```

Por lo cual usaremos las fichas descargando el .jsonld desde la web.

En limitaciones encontramos en primer lugar con que en el IAPH no hay información sobre el estado de conservación de cada bien, tampoco contempla horarios y tiene dos fallos de calidad: Latitud y longitud están intercambiadas. En la iglesia de la Encarnación de Albolote la marca en "latitud_s": -3.657 y "longitud_s": 37.23, y en cambio Albolote está en latitud 37.23 y longitud -3.66. Como vemos estos valores están puestos al revés. Cuando vayamos a usar las coordenadas tenemos que tener esto en cuenta. El segundo fallo de calidad que he encontrado es que la bibliografía a veces no corresponde ya que en este ejemplo aparece un libro sobre un yacimiento de la Edad del Bronce en Purullena, que no tiene nada que ver con esta iglesia. 
También descubrí otra limitación dentro de esta página, esta es que muchos registros dentro de Granada y provincia no cuentan con su municipio ni provincia, por lo que va a haber que identificarlos por el principio del código de los bienes de Granada (0118). 

La segunda fuente de datos que voy a usar es una dentro de los datos abiertos de la Junta de Andalucía, llamada Museos y Colecciones Museográficas de Andalucía: [Museos y Colecciones Museográficas de Andalucía](https://www.juntadeandalucia.es/datosabiertos/portal/dataset/museos-de-andalucia) dentro de ella para poder extraer nuestros datos necesarios referentes a los horarios de apertura, lugares en los que se encuentran vamos a introducirnos en el formato JSON: [JSON completo de museos](https://datos.juntadeandalucia.es/api/v0/museums/all?format=json) y al abrirla se descarga el archivo completo. 

Un ejemplo de extracción de datos con un museo sería el siguiente: 

[museo-ejemplo.json](docs/museo-ejemplo.json) 


Los campos que voy a usar aquí son sobre todo los horarios, días festivos abiertos u obviamente la localización de cada uno.

Como he podido observar, en el apartado del esfuerzo necesario para extraer los datos, aquí si va a haber que interpretarlos, ya que vienen como un texto libre y no como datos estructurados, por ejemplo en el opening hours aparece de martes a sábado en vez de cada día de la semana con su horario correspondiente, también el texto no es uniforme ya que repite el horario escrito de diferentes formas.
Además también las coordenadas tanto en longitud como en latitud aparecen con coma decimal, así que también habrá que reconvertirlas a número. 

Una de las limitaciones que encuentro es que los horarios y días de apertura solamente están para los 88 museos disponibles en la provincia de Granada, cosas que para los bienes culturales no, también es verdad que algunos de ellos no tienen horarios porque no están abiertos al público, solamente se pueden visitar desde fuera.

La tercera fuente que voy a usar para poder cumplimentar los horarios de visitas y sus respectivos días es añadir dos fuentes más con horarios de monumentos, para cubrir lo que no tienen ni el IAPH ni los museos.

  - [horarios del Patronato de la Alhambra](https://www.alhambra-patronato.es/visitar/horarios-y-tarifas)
  - [horario de visitas del Ayuntamiento de Granada](https://www.granada.org/inet/wagenda.nsf/byhac2/1D2DB77D7C9C7F8CC1258833003BED3D)

La primera fuente contiene los monumentos del Patronato, el conjunto de la Alhambra y los monumentos andalusíes (Corral del Carbón, Bañuelo, Casa Horno de Oro, Palacio de Dar al-Horra). 
La segunda fuente que es la del Ayuntamiento, el Carmen de los Mártires, el Palacio de los Córdova y Quinta Alegre. 

Las fuentes no proveen como las anteriores de datos abiertos, sino que son páginas web, así que tendré que leer el HTML y de ahí sacar la información, además los horarios vienen por temporadas (con rango de fechas). El programa tiene que mirar en qué fecha cae cada día del plan para saber qué horario aplicar, ya que en la página aparecen horarios distintos para cada temporada del año.

Dentro de las limitaciones he podido observar que solo hay horarios para algunos monumentos, la mayoría de la capital y casi nada de la provincia.


## Lógica de negocio

### Qué pide el usuario
El usuario le pide al programa un número de días, un horario establecido en la que el usuario pone a qué horas puede/quiere visitar, cuántas visitas quiere como máximo, cuál es su interés y la zona, indicada por municipios (por ejemplo, Granada capital o un municipio concreto de la provincia).

### Subproblema 1: calcular cuánto encaja cada sitio con el interés
Este paso prepara los datos para los siguientes; el núcleo de la lógica está en los subproblemas 2, 3 y 5.
El programa coge los datos de Granada del listado del IAPH, coge solo los sitios que empiezan por 0118, y de los museos coge la información del JSON, solo los de Granada también. Dependiendo de cada interés, el programa mira el periodo y el estilo de la ficha de cada bien y le asigna una puntuación según cuánto coincide con el interés. 
Para las etiquetas dentro de la guía digital del IAPH, buscamos por ejemplo Bañuelo y en sus etiquetas periodo y estilo aparece: 
Edad Media - Árabes y Arte islámico respectivamente y en json.ld son los campos periodos y den_estilo, que son los que usaremos. 
En los museos, el programa detecta en la descripción términos como nazarí, andalusí o islámico, y según eso les asigna una puntuación.
Estas fichas se descargan en .jsonld desde la Guía Digital.

### Subproblema 2: convertir los horarios a un formato común
Para poder descubrir cuándo visitar cada sitio, en los museos y los monumentos que tengan, se interpreta el texto para poder saber qué días abren y su horario. Los que no contemplen horario, se marcarán como horario desconocido. 
Los horarios vienen de tres formas distintas: texto libre en los museos, HTML por temporadas en el Patronato y el Ayuntamiento, y nada en la mayoría de los bienes. El programa los convierte todos a lo mismo: para cada sitio y cada día, si abre y de qué hora a qué hora, o en su caso desconocido si no tiene o no contempla este horario. Por ejemplo si pone que abre de martes a sábado de 10:00 a 19:00. Domingos y festivos de 10:00 a 15:00 -> da lunes cerrado, y abierto de martes a sábado de 10:00 a 19:00, y domingo de 10:00 a 15:00.

### Subproblema 3: construir el plan
Para repartir los días se colocan los sitios en los días que tiene el usuario, sin pasarse del máximo de visitas por día, sin repetirlos y poniendo los museos que estén abiertos. Se colocan solo cuando están abiertos dentro del rango horario del usuario.
Un sitio solo se coloca un día si está abierto ese día, según la temporada de esa fecha, y dentro de las horas que el usuario ha dicho. Esto vale para museos y monumentos.

### Subproblema 4: sitios sin horario
Esto hace referencia para los que están marcados como "horario desconocido", van a estar marcados en cualquier día como "visitables desde fuera" y solo entrarán si sobra hueco después de los que tienen horario, avisando además al usuario de que no tiene horario oficial establecido.

### Subproblema 5: qué hacer si no cabe todo
Si no entra todo en los días disponibles, se usa la puntuación calculada en el subproblema 1: primero los que coinciden en periodo y estilo, luego los que solo coinciden en uno, y después los museos que lo mencionan en la descripción.

### Cómo se comprueba que es correcto
Para comprobar que el plan es correcto, podemos verificar que ningún día pasa del máximo, todo encaja con el interés, no se repite nada y tanto museos como monumentos aparecen los días que abren.
Además de comprobar que ningún sitio está fuera de las horas del usuario, que el horario de ejemplo se convierte bien y que los sitios sin horario siguen la regla del subproblema 4.




## Juego de rol

![Ficha del juego de rol](docs/ficha.jpg)

## Documentación

- [Configuración del entorno](docs/configuracion.md)
- [Personas](docs/personas-culturagrx.md)
- [User journeys](docs/user-journeys-culturagrx.md)
- [Historias de usuario](docs/historias-de-usuario.md)
- [Milestones](docs/milestones.md)
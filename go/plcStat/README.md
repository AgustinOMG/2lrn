# plcStat 

La funcionalidad principal de este programa es la obtencion de los parametros de de CPU y memoria del PLC,
"Se considera que el PLC corre una version de LINUX, las pruebas seran realizadas en un PLC de la marca WAGO , Modelo 750 - 8206"

Una vez obtenido los datos del PLC se generara un servidor HTTP el cual servira los datos de estado del PLC a cualqueir sistema capaz de realizar uan solicitud de tipo GET.

## LOG 

* se genera el packete
* se selecciona la libreria
* se prueba la obtencion de de datos de CPU y memoria  -  27/08/2022
* se copia la infromacion de la biblioteca exclusivamente para linux con el fin de reducir el binario.
* se modifican los archivos para tarbajar con nuetsra nueva estructura.
* se prueba la operacion, y se compila para plc WAGO PFC200 750-8206 ARMV8 64 Bits
* Se optimiza el binario removiendo la infroamcion del debbuging 8/09/2022



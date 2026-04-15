# Reflexion - Calculadora en GO
1.	¿Cuántas líneas tiene tu función main al final del taller? Cuéntalas con cuidado.
   Mi funcion main tiene aproximadamente 75 lineas de codigo.
  	Esto ocurre ya que toda la logica del programa esta concentrada en una sola funcion.
  	
2.	Si tuvieras que agregar 5 operaciones más (raíz cuadrada, logaritmo, seno, coseno, módulo),
   ¿qué tan grande se haría tu main? ¿Sería fácil de leer para alguien que vea el código por primera vez?
  	La funcion main creceria bastante, aproximadamente entre unas 130 o 150 lineas. No seria facil de leer porque el swictch tendria muchos mas cosas y el codigo se volveria mas largo y dificil de entender.
  	
3.	Notaste que el código para 'pedir un número al usuario' o 'imprimir el resultado' se repite varias veces.
   ¿No sería mejor escribirlo una sola vez y reutilizarlo en muchos lugares?
  	En mi codigo se repite el proceso de pedir datos al usuario y manejar resulrado. Seria mejor crear funciones para evitar repetir codigo, mejorar el orden y facilitar posibles cambios.
  	
4.	Tu historial es una variable string gigante.
   ¿Qué pasaría si quisieras: ordenarlo alfabéticamente, eliminar la operación número 2, o contar cuántas veces se usó la operación de suma?
  	Seria dificil de manejar porque es solo un string. No se podria modificar facilmente. Lo ideal seria usar un arreglo dinamico.
  	
5.	Después de este taller, ¿qué fue lo más difícil de Go para ti? ¿Y lo más interesante?
   Lo mas dificil fue manejar los tipos de datos y conversiones. Los mas interesante fue usar switch y for para resolver las operaciones.

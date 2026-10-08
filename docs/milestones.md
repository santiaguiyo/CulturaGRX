# Milestones

Los dos milestones avanzan sobre el mismo problema, el de la HU001, que es el que hace referencia a los horarios difíciles de entender.

## Milestone 0: Primera versión del problema en código

**Objetivo:** analizar el problema de la HU001 siguiendo el diseño dirigido por el dominio y pasarlo a código. Es un producto interno: todavía no tiene lógica de negocio.

**Entrega:**
El código, en el lenguaje elegido, en el que cada parte sale de un issue. Cada issue plantea un problema de la HU001, y en él se decide si ese concepto es un objeto valor, una entidad o un agregado, y por qué.

**Comprobación:**
Para comprobar que el milestone está bien realizado, todos los issues deben partir de la HU001 o de alguno de los elementos relacionados con ella, como el user journey o las fuentes de datos. Además, cada issue debe plantear un problema concreto, y no una tarea, y explicar si el concepto correspondiente se implementa como objeto valor, entidad o agregado, junto con la justificación de esa decisión.

Cada parte del código debe resolver uno de esos issues y los commits tienen que indicar claramente qué issue solucionan. También se comprobará que el código sea sintácticamente correcto y que mantenga la estructura habitual de un proyecto desarrollado en el lenguaje elegido. Por último, los cambios deberán haber sido revisados y aprobados mediante un pull request.

## Milestone 1: Primer caso completo resuelto

**Objetivo:** programar, sobre lo entregado en el milestone 0, la lógica de negocio que resuelve un caso sencillo del problema de la HU001, dividiéndolo en problemas más pequeños que se puedan comprobar con un test. Sigue siendo interno, todavía sin interfaz para el usuario.

**Entrega:**
La lógica de negocio y sus tests, incluidos los de los errores que puedan darse.

**Comprobación:**
Para comprobar que el milestone es válido, los tests deben poder lanzarse con una única orden del gestor de tareas y pasar todos. Los tests tienen que cubrir tanto el caso sencillo planteado como los posibles errores que puedan aparecer.

Además, cada test debe estar relacionado con un issue concreto, y cada uno de esos issues debe estar relacionado con la HU001. Como en el milestone anterior, los cambios realizados tendrán que haber sido revisados y aprobados mediante un pull request.
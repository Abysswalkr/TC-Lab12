# TC-Lab12 — README

Proyecto con **4 problemas** de Teoría de la Computación / Programación funcional.
Los componentes finales (código) para los problemas 2–4 están en **Go** y emplean **funciones anónimas (lambdas)**. El Problema 1 es teórico (cálculo-λ).

---

## Estructura

```
.
├── problema2.go   # Ordenar lista de “diccionarios” (map[string]any) por key
├── problema3.go   # Transpuesta de una matriz (genérico con generics)
├── problema4.go   # Eliminar elementos indicados de una lista
└── README.md
```

> Requisitos: Go 1.20+
> Ejecución: `go run problema2.go`, `go run problema3.go`, `go run problema4.go`

---

## Video de ejecución

Puedes ver la ejecución completa del código aquí: **[https://youtu.be/9WYvdApjNq0](https://youtu.be/9WYvdApjNq0)**

---

## Problema 1 (25%)

**Enunciado:**

1. Escribir la **reducción-β** de `NOT`.
2. Explicar cómo se ven **recursión** y **ciclos**.
3. Explicar **cuándo usar** este estilo y **cuándo no**, con ejemplo.

**Solución (resumen):**

* **Booleans de Church:**
  `TRUE ≡ λx.λy.x`, `FALSE ≡ λx.λy.y`, `NOT ≡ λp. p FALSE TRUE`.

* **Reducciones-β principales:**

  * `NOT TRUE`
    `(λp. p F T) TRUE → TRUE F T → (λx.λy.x) F T → (λy.F) T → F`.
  * `NOT FALSE`
    `(λp. p F T) FALSE → FALSE F T → (λx.λy.y) F T → (λy.y) T → T`.

* **Recursión:** con el combinador de punto fijo **Y**
  `Y ≡ λf.(λx.f (x x)) (λx.f (x x))`, cumple `Y F = F (Y F)`.
  Ejemplo (esqueleto): `FACT ≡ Y (λrec.λn. IF (ISZERO n) 1 (MUL n (rec (PRED n))))`.

* **Ciclos:** se modelan como **recursión de cola** o con **folds**.
  `WHILE cond step state = IF (cond state) (WHILE cond step (step state)) state`.

* **Cuándo usar / no usar (breve):**

  * **Usar**: transformación de datos inmutables, composición, concurrencia sin estado compartido (ej.: `filter`→`map`→`sort`).
  * **Evitar/mitigar**: algoritmos que requieren **mutación in-place** intensiva o recursión profunda sin TCO (ej.: DP en matrices gigantes, Floyd–Warshall); I/O muy stateful.

---

## Problema 2 (25%) — Ordenar lista de diccionarios por `key` con lambda (Go)

**Idea:** ordenar `[]map[string]any]` por la llave dada usando `sort.SliceStable` con **comparador lambda**.
Los elementos sin la key se envían al final.

**Cómo ejecutar**

```bash
go run problema2.go
```

**Salida obtenida**

```
map[color:Gold make:Mi Max model:2]
map[color:Blue make:Samsung model:7]
map[color:Black make:Nokia model:216]
```

**Interpretación:** orden ascendente por `model`: `2 < 7 < 216`.

---

## Problema 3 (25%) — Transpuesta de una matriz con lambda (Go)

**Idea:** función genérica `Transpose[T any]` que valida matriz rectangular y usa una **lambda** `makeRow := func(j int) []T { ... }` para construir cada fila de `Xᵀ`.

**Cómo ejecutar**

```bash
go run problema3.go
```

**Salida obtenida**

```
[1 4 7]
[2 5 8]
[3 6 9]
```

**Interpretación:** `Xᵀ` de la matriz 3×3 del enunciado.

---

## Problema 4 (25%) — Eliminar elementos indicados usando lambda (Go)

**Idea:** crear un `set` (map) con los elementos a borrar y un **predicado lambda** `keep := func(x T) bool { ... }` dentro de `RemoveElements[T comparable]` para filtrar preservando orden.

**Cómo ejecutar**

```bash
go run problema4.go
```

**Salida obtenida**

```
[rojo verde azul gris negro]
```

**Interpretación:** se eliminaron `amarillo` y `blanco`; `café` se ignoró por no existir en la lista original.

---

## Notas técnicas

* Uso de **lambdas** en Go:

  * `sort.SliceStable(..., func(i, j int) bool { ... })` (Prob. 2)
  * `makeRow := func(j int) []T { ... }` (Prob. 3)
  * `keep := func(x T) bool { ... }` (Prob. 4)
* Se preserva **inmutabilidad lógica** (se retornan nuevas colecciones).
* Los ejemplos imprimen la salida esperada para verificación rápida.

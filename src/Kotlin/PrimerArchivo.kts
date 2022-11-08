/**
 * You can edit, run, and share this code.
 * play.kotlinlang.org
 */
fun main() {
    val nombre = "Jose"
    val apellido: String = "Carrasco"
    println("${10 + 3}")
    println("hola $nombre $apellido, como esta?")
//     Long = 9220000000000000000
    var numero: Long = 100
    println("${
            	numero*(numero+1)/2
            }")
//     for (n in 1..10){
//         print("$n ")
//     }
//     println()
    
//     while(numero>0){
//         println("$numero")
//         numero--
//     }
//     
//     Sieve(100)
//     val apellido2: String?
//    	val numero: Int = 2
//     if{
//         appellido2?isNullorEmpty()
//     }
    val persona = Persona("Jose","Carrasco")
    
    println("${persona.Nombre}")
    
    var arregloNuevo = arrayOf<Int>(100)
//     companion object{
//         val ejemplo = Ejemplo(
//             "un",
//             "ejemplo",
//             0
//         )
// 	}
    
    println(saludar("Jose Carrasco"))
    
}

fun Sieve(primo: Int): Array<Int>{
    //     Sieve
	var arreglo = arrayOf<Int>(primo+1)
    var i: Int = 0
    while(i<=primo){
       arreglo[i]=i
       i++
    }
    arreglo[1]=0
    i = 2
    while(i*i<=primo){
        if(arreglo[i]!=0){
            var j: Int = i*i
            while(j<=primo){
                arreglo[j]=0
                j+=i
            }   
        }
        i++
    }
    return arreglo
}

// fun saludar(nuevoNombre: String): String{
// 	return "Hola $nuevoNombre"    
// }

fun saludar(nuevoNombre: String) = "hola $nuevoNombre"

data class Persona(
	var Nombre: String,
    var Apellido: String
)
/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/30 22:38
 * @version     v1.0
 * @filename    CustomSorting.go
 * @description
 ***************************************************************************/
package main

func customSorting(strArr []string) []string {
	var odd []string
	var even []string
	for i := 0; i < len(strArr); i++ {
		if len(strArr[i])%2 == 1 {
			odd = append(odd, strArr[i])
		}
		if len(strArr[i])%2 == 0 {
			even = append(even, strArr[i])
		}
	}
	odd = sort1(odd, 0)
	even = sort1(even, 1)
	return append(odd, even...)
}

func sort1(arr []string, mode int) []string {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			switch mode {
			case 0:
				if len(arr[i]) > len(arr[j]) {
					arr[i], arr[j] = arr[j], arr[i]
				}
				if len(arr[i]) == len(arr[j]) {
					if arr[i] > arr[j] {
						arr[i], arr[j] = arr[j], arr[i]
					}
				}
			case 1:
				if len(arr[i]) < len(arr[j]) {
					arr[i], arr[j] = arr[j], arr[i]
				}
				if len(arr[i]) == len(arr[j]) {
					if arr[i] > arr[j] {
						arr[i], arr[j] = arr[j], arr[i]
					}
				}
			}
		}
	}
	return arr
}

//func main() {
//	arr := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
//	fmt.Println(customSorting(arr))
//}

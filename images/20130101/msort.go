package main

import (
	"fmt"
)

// 冒泡排序
func bubble(arr []int) (narr []int) {
	larr := len(arr)
	for i := 0; i < larr; i++ {
		for j := 0; j < larr-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// 选择排序
func selection(arr []int) (narr []int) {
	larr := len(arr)
	for i := 0; i < larr; i++ {
		minidx := i
		for j := i + 1; j < larr; j++ {
			if arr[j] < arr[minidx] {
				minidx = j
			}
		}
		arr[i], arr[minidx] = arr[minidx], arr[i]
	}

	return arr
}

// 插入排序
func insertion(arr []int) (narr []int) {
	larr := len(arr)
	for i := 1; i < larr; i++ {
		preindex, current := i-1, arr[i]
		for ; preindex >= 0 && arr[preindex] > current; preindex-- {
			arr[preindex+1] = arr[preindex]
		}
		arr[preindex+1] = current
	}

	return arr
}

// 希尔排序
func shell(arr []int) (narr []int) {
	return arr
}

// 归并排序
func merge_sort(arr []int, left, mid, right int) (narr []int) {
	i, j := left, mid+1
	temp := []int{}
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}

	for j <= right {
		temp = append(temp, arr[j])
		j++
	}

	for l := 0; l < len(temp); l++ {
		arr[left+l] = temp[l]
	}

	return arr
}

func merge_conquer(arr []int, left, right int) (narr []int) {
	if left >= right {
		return arr
	}

	mid := (left + right) / 2
	merge_conquer(arr, left, mid)
	merge_conquer(arr, mid+1, right)
	return merge_sort(arr, left, mid, right)
}

func merge(arr []int) (narr []int) {
	return merge_conquer(arr, 0, len(arr)-1)
}

// 快速排序
func quick_partition(arr []int, left, right int) (loc int) {
	pivot := arr[left]
	for left < right {
		for ; left < right && arr[right] > pivot; right-- {
		}
		arr[left] = arr[right]
		for ; left < right && arr[left] <= pivot; left++ {
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func quick_sort(arr []int, left, right int) {
	if !(left < right && left >= 0) {
		return
	}

	loc := quick_partition(arr, left, right)
	quick_sort(arr, left, loc-1)
	quick_sort(arr, loc+1, right)
}

func quick(arr []int) (narr []int) {
	quick_sort(arr, 0, len(arr)-1)
	return arr
}

// 堆排序
// 最大堆：最大堆中的最大元素值出现在根结点（堆顶），堆中每个父节点的元素值都大于等于其孩子结点（如果存在）
// 堆公式：
//     Parent(i) = floor((i-1)/2)，i 的父节点下标
//     Left(i) = 2i + 1，i 的左子节点下标
//     Right(i) = 2(i + 1)，i 的右子节点下标
// 堆排序：堆排序就是把最大堆堆顶的最大数取出，将剩余的堆继续调整为最大堆，再次将堆顶的最大数取出，这个过程持续到剩余数只有一个时结束。
//     在堆中定义以下几种操作：
//       最大堆调整（Heapify）：将堆的末端子节点作调整，使得子节点永远小于父节点
//       创建最大堆（Build）：将堆所有数据重新排序，使其成为最大堆
//       堆排序（Heap-Sort）：移除位在第一个数据的根节点，并做最大堆调整的递归运算
func heap_build(arr []int, hsize int) (narr []int) {
	for i := hsize / 2; i >= 0; i-- {
		heap_heapify(arr, i, hsize)
	}
	return arr
}

func heap_heapify(arr []int, idx, hsize int) {
	left, right, max := 2*idx+1, 2*idx+2, idx
	if left < hsize && arr[left] > arr[max] {
		max = left
	}
	if right < hsize && arr[right] > arr[max] {
		max = right
	}
	if max != idx {
		arr[idx], arr[max] = arr[max], arr[idx]
		heap_heapify(arr, max, hsize)
	}
}

func heap(arr []int) (narr []int) {
	hsize := len(arr)
	heap_build(arr, hsize)
	for i := hsize - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heap_heapify(arr, 0, i)
	}
	return arr
}

// main测试入口
func main() {
	iarr := []int{7, 1, 5, 9, 5, 11, 3}
	// TODO: 选择排序方法
	oarr := heap(iarr)
	for i := 0; i < len(oarr); i++ {
		fmt.Printf("%d ", oarr[i])
	}
	fmt.Println()
}

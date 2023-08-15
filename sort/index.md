<style></style>

<div class="markdown prose w-full break-words dark:prose-invert light"><p>当然可以，以下是几种常见的排序算法的代码示例、时间复杂度解释和简要介绍。</p><p><strong>1. 冒泡排序：</strong></p><p>代码示例：</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>golang</span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>Copy code</button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-golang"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">bubbleSort</span><span class="hljs-params">(arr []<span class="hljs-type">int</span>)</span></span> {
    n := <span class="hljs-built_in">len</span>(arr)
    <span class="hljs-keyword">for</span> i := <span class="hljs-number">0</span>; i &lt; n<span class="hljs-number">-1</span>; i++ {
        <span class="hljs-keyword">for</span> j := <span class="hljs-number">0</span>; j &lt; n-i<span class="hljs-number">-1</span>; j++ {
            <span class="hljs-keyword">if</span> arr[j] &gt; arr[j+<span class="hljs-number">1</span>] {
                arr[j], arr[j+<span class="hljs-number">1</span>] = arr[j+<span class="hljs-number">1</span>], arr[j]
            }
        }
    }
}
</code></div></div></pre><p>时间复杂度解释：
冒泡排序的平均时间复杂度是 O(n^2)，其中 n 是数组的大小。由于需要比较相邻元素并交换位置，因此在最坏情况下需要进行 n*(n-1)/2 次比较和交换。</p><p><strong>2. 插入排序：</strong></p><p>代码示例：</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>golang</span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>Copy code</button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-golang"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">insertionSort</span><span class="hljs-params">(arr []<span class="hljs-type">int</span>)</span></span> {
    n := <span class="hljs-built_in">len</span>(arr)
    <span class="hljs-keyword">for</span> i := <span class="hljs-number">1</span>; i &lt; n; i++ {
        key := arr[i]
        j := i - <span class="hljs-number">1</span>
        <span class="hljs-keyword">for</span> j &gt;= <span class="hljs-number">0</span> &amp;&amp; arr[j] &gt; key {
            arr[j+<span class="hljs-number">1</span>] = arr[j]
            j--
        }
        arr[j+<span class="hljs-number">1</span>] = key
    }
}
</code></div></div></pre><p>时间复杂度解释：
插入排序的平均时间复杂度是 O(n^2)，其中 n 是数组的大小。在最坏情况下，当数组是逆序时，需要进行 n*(n-1)/2 次比较和移动。</p><p><strong>3. 归并排序：</strong></p><p>代码示例：</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>golang</span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>Copy code</button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-golang"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">mergeSort</span><span class="hljs-params">(arr []<span class="hljs-type">int</span>)</span></span> []<span class="hljs-type">int</span> {
    <span class="hljs-keyword">if</span> <span class="hljs-built_in">len</span>(arr) &lt;= <span class="hljs-number">1</span> {
        <span class="hljs-keyword">return</span> arr
    }

    mid := <span class="hljs-built_in">len</span>(arr) / <span class="hljs-number">2</span>
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    
    <span class="hljs-keyword">return</span> merge(left, right)
}

<span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">merge</span><span class="hljs-params">(left, right []<span class="hljs-type">int</span>)</span></span> []<span class="hljs-type">int</span> {
result := <span class="hljs-built_in">make</span>([]<span class="hljs-type">int</span>, <span class="hljs-number">0</span>, <span class="hljs-built_in">len</span>(left)+<span class="hljs-built_in">len</span>(right))
i, j := <span class="hljs-number">0</span>, <span class="hljs-number">0</span>
<span class="hljs-keyword">for</span> i &lt; <span class="hljs-built_in">len</span>(left) &amp;&amp; j &lt; <span class="hljs-built_in">len</span>(right) {
<span class="hljs-keyword">if</span> left[i] &lt; right[j] {
result = <span class="hljs-built_in">append</span>(result, left[i])
i++
} <span class="hljs-keyword">else</span> {
result = <span class="hljs-built_in">append</span>(result, right[j])
j++
}
}
result = <span class="hljs-built_in">append</span>(result, left[i:]...)
result = <span class="hljs-built_in">append</span>(result, right[j:]...)
<span class="hljs-keyword">return</span> result
}
</code></div></div></pre><p>时间复杂度解释：
归并排序的平均时间复杂度是 O(n log n)，其中 n 是数组的大小。它采用分治法，将数组分成两个子数组，分别排序后再合并。</p><p><strong>4. 选择排序：</strong></p><p>代码示例：</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>golang</span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>Copy code</button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-golang"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">selectionSort</span><span class="hljs-params">(arr []<span class="hljs-type">int</span>)</span></span> {
n := <span class="hljs-built_in">len</span>(arr)
<span class="hljs-keyword">for</span> i := <span class="hljs-number">0</span>; i &lt; n<span class="hljs-number">-1</span>; i++ {
minIndex := i
<span class="hljs-keyword">for</span> j := i + <span class="hljs-number">1</span>; j &lt; n; j++ {
<span class="hljs-keyword">if</span> arr[j] &lt; arr[minIndex] {
minIndex = j
}
}
arr[i], arr[minIndex] = arr[minIndex], arr[i]
}
}
</code></div></div></pre><p>时间复杂度解释：
选择排序的平均时间复杂度是 O(n^2)，其中 n 是数组的大小。每次在未排序部分找到最小值，并将其交换到已排序部分的末尾。</p><p><strong>5. 希尔排序：</strong></p><p>代码示例：</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>golang</span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg>Copy code</button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-golang"><span class="hljs-function"><span class="hljs-keyword">func</span> <span class="hljs-title">shellSort</span><span class="hljs-params">(arr []<span class="hljs-type">int</span>)</span></span> {
n := <span class="hljs-built_in">len</span>(arr)
gap := n / <span class="hljs-number">2</span>
<span class="hljs-keyword">for</span> gap &gt; <span class="hljs-number">0</span> {
<span class="hljs-keyword">for</span> i := gap; i &lt; n; i++ {
temp := arr[i]
j := i
<span class="hljs-keyword">for</span> j &gt;= gap &amp;&amp; arr[j-gap] &gt; temp {
arr[j] = arr[j-gap]
j -= gap
}
arr[j] = temp
}
gap /= <span class="hljs-number">2</span>
}
}
</code></div></div></pre><p>时间复杂度解释：
希尔排序的时间复杂度取决于增量序列的选择。最好的情况下是 O(n log n)，但最坏情况下可能接近 O(n^2)。希尔排序是插入排序的改进版本，通过逐步减小步长来对数据进行分组和排序。</p><p>这些排序算法各有特点，适用于不同的场景。选择合适的排序算法取决于输入数据的规模和特性。</p></div>
package sort.insertion;

import java.util.Arrays;

public class Insertion {

    public static void main(String[] args) {
        int[] unsortedItems = {5, 2, 9, 41, -6, 58, 12, 6};
        sort(unsortedItems);
        Arrays.stream(unsortedItems).forEach(System.out::println);
    }

    public static void sort(int[] arr){
        int key, j;
        for (int i = 1; i < arr.length; i++) {
            key = arr[i];
            j = i-1;
            while (j >= 0 && arr[j] > key){
                arr[j+1] = arr[j];
                j = j-1;
            }
            arr[j + 1] = key;
        }
    }
}

package sort.selection;


import java.util.Arrays;

public class Selection {

    public static void main(String[] args) {

//        Comparable[] unsortedItems = {5, 2, 9, 41, -6, 58, 12, 6};

//        Comparable[] unsortedItems = {5.5, 2.5, 9.0, 41.25, -6.0, 58.01, 12.2, 6.9};

        Comparable[]  unsortedItems = {"Longing" , "rusted", "seventeen", "daybreak", "furnace", "nine", "benign", "homecoming", "one", "freight car"};
        sort(unsortedItems);
        Arrays.stream(unsortedItems).forEach(System.out::println);
    }

    public static void sort(Comparable[] arr){
        int n = arr.length;
        int min;
        for (int i = 0; i < n - 1; i++) {
            min = i;
            for (int j = i+1; j < n; j++) {
                if (arr[j].compareTo(arr[min]) < 0) min = j;
            }
//            arr[i] = arr[i] ^ arr[min] ^ (arr[min] = arr[i]); // the statement handles swapping when we use int[] in place of Comparable[]
            Comparable temp = arr[i];
            arr[i] = arr[min];
            arr[min] = temp;
        }
    }
}

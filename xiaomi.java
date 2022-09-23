import java.text.BreakIterator;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.List;

class xiaomi {
    public static void main(String[] args) {
        int n;
        Scanner sc = new Scanner(System.in);
        n = sc.nextInt();
        List<Integer> nums = new ArrayList<Integer>();
        for (int i = 0; i< n; i++) {
            nums.add(i+1);
        }
        int three = 0;
        int total = 0;
        for (int i = 0; i < nums.size(); i = (i+1)%nums.size()) {
            if (nums.get(i) > 0) {
                three++;
            }
            if (three == 3) {
                nums.set(i, 0-nums.get(i));
                total++;
                three = 0;
            }
            if (total == nums.size()-1) {
                break;
            }
        }
        nums.forEach(num -> {
            if (num > 0) {
                System.out.println(num);
                return;
            }
        });
    }
    }
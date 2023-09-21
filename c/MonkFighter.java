package c;

import java.util.Random;

public class MonkFighter implements Comparable {
    private Integer chi_energy;
    private String monastery;
    private int monkFighterId;
    private static int countMonkFighters = 0;
    public MonkFighter() {
        Random rand = new Random();
        chi_energy = rand.nextInt(100);
        monastery = (rand.nextInt(2) == 0) ? "Guan-Ying" : "Guan-Yang";
        monkFighterId = countMonkFighters++;
    }

    @Override
    public String toString() {
        return "Monk Fighter Number " + monkFighterId + " (from " + monastery + " monastery) | Chi Energy: " + chi_energy;
    }

    public int compareTo(Object obj) {
        MonkFighter other = (MonkFighter) obj;
        if(this.chi_energy > other.chi_energy) return 1;
        else if(this.chi_energy < other.chi_energy) return -1;
        else return 0;
    }
    static MonkFighter FIGHT(MonkFighter left, MonkFighter right){
        if(left.chi_energy > right.chi_energy) return left;
        else return right;
    }
}

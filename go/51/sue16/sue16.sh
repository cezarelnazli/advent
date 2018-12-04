grep "children: 3" sue16.in >> sue16.out
grep "cats: 7" sue16.in >> sue16.out
grep "samoyeds: 2" sue16.in >> sue16.out
grep "akitas: 0" sue16.in >> sue16.out
grep "vizslas: 0" sue16.in >> sue16.out
grep "goldfish: 5" sue16.in >> sue16.out
grep "trees: 3" sue16.in >> sue16.out
grep "cars: 2" sue16.in >> sue16.out
grep "perfumes: 1" sue16.in >> sue16.out

sort sue16.out | uniq > sue16.out.1

# string of grep -v's

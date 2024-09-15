def main():
    all_states = set(["MT", "WA", "OR", "ID", "NV", "UT", "CA", "AZ"])

    stations = {}
    stations["k_one"] = set(["ID", "NV", "UT"])
    stations["k_two"] = set(["WA", "ID", "MT"])
    stations["k_three"] = set(["OR", "NV", "CA"])
    stations["k_four"] = set([ "NV", "UT"])
    stations["k_five"] = set([ "CA", "AZ"])

    final_stations = set()

    while all_states:
        better_station = None
        covered_states = set()

        for station, states in stations.items():
            covered = all_states & states
            if len(covered) > len(covered_states):
                better_station = station
                covered_states = covered

        all_states -= covered_states
        final_stations.add(better_station)

    print("Final stations:", final_stations)

main()

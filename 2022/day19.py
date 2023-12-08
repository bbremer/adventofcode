from concurrent.futures import ProcessPoolExecutor
import re

from input_parser import parse_day


def num_geodes(args):
    int_groups, budget = args

    ore_per_orerobot, ore_per_clayrobot, *int_groups = int_groups
    ore_per_obsrobot, clay_per_obsrobot, *int_groups = int_groups
    ore_per_geoderobot, obs_per_geoderobot = int_groups

    max_ore_needed = max(ore_per_orerobot, ore_per_clayrobot, ore_per_obsrobot,
                         ore_per_geoderobot)

    memo = {}

    def iteration(budget, ore_bots, clay_bots, obs_bots, ore, clay, obs):
        if budget == 1:
            return 0

        state = (budget, ore_bots, clay_bots, obs_bots, ore, clay, obs)
        if state in memo:
            return memo[state]

        if ore >= ore_per_geoderobot and obs >= obs_per_geoderobot:
            new_ore = ore - ore_per_geoderobot
            new_obs = obs - obs_per_geoderobot
            max_geodes = iteration(budget-1,
                                   ore_bots, clay_bots, obs_bots,
                                   new_ore + ore_bots, clay + clay_bots,
                                   new_obs + obs_bots)
            max_geodes += budget - 1
            memo[state] = max_geodes
            return max_geodes

        max_geodes = iteration(budget-1,
                               ore_bots, clay_bots, obs_bots,
                               ore + ore_bots, clay + clay_bots,
                               obs + obs_bots)

        if ore >= ore_per_orerobot and ore_bots < max_ore_needed:
            new_ore = ore - ore_per_orerobot
            new_geodes = iteration(budget-1,
                                   ore_bots+1, clay_bots, obs_bots,
                                   new_ore + ore_bots, clay + clay_bots,
                                   obs + obs_bots)
            if new_geodes > max_geodes:
                max_geodes = new_geodes

        if ore >= ore_per_clayrobot and clay_bots < clay_per_obsrobot:
            new_ore = ore - ore_per_clayrobot
            new_geodes = iteration(budget-1,
                                   ore_bots, clay_bots+1, obs_bots,
                                   new_ore + ore_bots, clay + clay_bots,
                                   obs + obs_bots,)
            if new_geodes > max_geodes:
                max_geodes = new_geodes

        if (ore >= ore_per_obsrobot and clay >= clay_per_obsrobot
                and obs_bots < obs_per_geoderobot):
            new_ore = ore - ore_per_obsrobot
            new_clay = clay - clay_per_obsrobot
            new_geodes = iteration(budget-1,
                                   ore_bots, clay_bots, obs_bots+1,
                                   new_ore + ore_bots,
                                   new_clay + clay_bots,
                                   obs + obs_bots)
            if new_geodes > max_geodes:
                max_geodes = new_geodes

        memo[state] = max_geodes

        return max_geodes

    return iteration(budget, 1, 0, 0, 0, 0, 0)


if __name__ == '__main__':
    test_data = ("Blueprint 1:"
                 " Each ore robot costs 4 ore."
                 " Each clay robot costs 2 ore."
                 " Each obsidian robot costs 3 ore and 14 clay."
                 " Each geode robot costs 2 ore and 7 obsidian."
                 "\n"
                 "Blueprint 2:"
                 " Each ore robot costs 2 ore."
                 " Each clay robot costs 3 ore."
                 " Each obsidian robot costs 3 ore and 8 clay."
                 " Each geode robot costs 3 ore and 12 obsidian.")

    # input_iter = test_data.split('\n')
    input_iter = parse_day(19)

    template = re.compile(
            "Blueprint ([0-9]+):"
            " Each ore robot costs ([0-9]+) ore."
            " Each clay robot costs ([0-9]+) ore."
            " Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay."
            " Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian."
        )

    inputs = []
    for blueprint in input_iter:
        match = template.match(blueprint)
        bp_id, *int_groups = [int(x) for x in match.groups()]
        inputs.append(int_groups)

    with ProcessPoolExecutor() as executor:
        p1_inputs = ((i, 24) for i in inputs)
        nums_iter = list(executor.map(num_geodes, p1_inputs))
        nums = []
        for i, n in enumerate(nums_iter, 1):
            nums.append(n)
        print("Part 1:", sum(i*n for i, n in enumerate(nums, 1)))

        p2_inputs = ((i, 32) for i in inputs[:3])
        n1, n2, n3 = list(executor.map(num_geodes, p2_inputs))
        print("Part 2:", n1*n2*n3)

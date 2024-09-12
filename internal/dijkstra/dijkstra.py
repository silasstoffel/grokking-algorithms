# Graph:
#       ---6-> A ---1-----
#      |       /\         |
#  start       || (3)     end
#      |       ||         |
#        ---2-> B ---5----
#

# {
# 'a': 'b',
# 'b': 'start',
# 'end': 'a'
#}
# start -> b -> a -> end
graph = {}

graph["start"] = {}
graph["start"]["a"] = 6
graph["start"]["b"] = 2

graph["a"] = {}
graph["a"]["end"] = 1

graph["b"] = {}
graph["b"]["a"] = 3
graph["b"]["end"] = 5

graph["end"] = {}

infinite = float("inf")
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["end"] = infinite

parents = {}
parents["a"] = "start"
parents["b"] = "start"
parents["end"] = None

processed = []

def find_lowest_cost_node(costs):
    lowest_cost = float("inf")
    lowest_cost_node = None

    for node in costs:
        if node not in processed and costs[node] < lowest_cost:
            lowest_cost = costs[node]
            lowest_cost_node = node

    return lowest_cost_node

def print_path(parents, start_key, end_key):

    if parents.get(end_key) is None:
        return None

    items = [end_key]
    key = end_key

    max_keys = len(parents.keys())
    verification_counter = 0

    while verification_counter <= max_keys:
        item = parents[key]

        if item == start_key:
            items.append(item)
            break

        items.append(item)
        key = item
        verification_counter += 1

    sorted_items =  items[::-1]
    print(' -> '.join(sorted_items))


node = find_lowest_cost_node(costs)
while node is not None:
    cost = costs[node]
    neighbors = graph[node]

    for neighbor in neighbors.keys():
        new_cost = cost + neighbors[neighbor]

        if costs[neighbor] > new_cost:
            costs[neighbor] = new_cost
            parents[neighbor] = node

    processed.append(node)
    node = find_lowest_cost_node(costs)


print_path(parents, 'start', 'end')

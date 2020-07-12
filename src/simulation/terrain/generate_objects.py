import os
import argparse
import csv 

import overpy 


try: 
    parser = argparse.ArgumentParser() 
    parser.add_argument('-wdir', type=str, default=r'/home/miguel/Documents/projects/Wildfire', help="working dir")
    parser.add_argument('lat1', type=float, help="latitude1")
    parser.add_argument('lon1', type=float, help="longitude1")
    parser.add_argument('lat2', type=float, help="latitude2")
    parser.add_argument('lon2', type=float, help="longitude2")
 
    args, unknown = parser.parse_known_args()

    os.chdir(args.wdir) 
except Exception as e:
    print(e)
lat1 = args.lat1
lon1 = args.lon1
lat2 = args.lat2
lon2 = args.lon2
# print(lat1,lat2, lon1, lon2)
 
api = overpy.Overpass()

q = f"""
    way({lat1}, {lon1}, {lat2}, {lon2}) ["highway"];
    (._;>;);
    out body;
    """ 
result = api.query(q) 

street_lst = list()
for way in result.ways: 
    for node in way.nodes:
        street_lst.append((float(node.lat), float(node.lon))) 
  
with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/street_coordinates.csv', "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerows(street_lst)


api = overpy.Overpass()
q = f"""
    way({lat1}, {lon1}, {lat2}, {lon2}) ["landuse"];
    (._;>;);
    out body;
    """ 
result = api.query(q) 

natural_lst = list()
for way in result.ways: 
    for node in way.nodes:
        natural_lst.append((float(node.lat), float(node.lon))) 
  
with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/landuse_coordinates.csv', "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerows(natural_lst)


api = overpy.Overpass()
q = f"""
    way({lat1}, {lon1}, {lat2}, {lon2}) ["natural"];
    (._;>;);
    out body;
    """ 
result = api.query(q) 

natural_lst = list()
for way in result.ways: 
    for node in way.nodes:
        natural_lst.append((float(node.lat), float(node.lon))) 
  
with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/natural_coordinates.csv', "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerows(natural_lst)
 
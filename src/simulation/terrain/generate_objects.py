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



def extract_elements(lat1, lat2, lon1, lon2, tag):
    api = overpy.Overpass()

    q = f"""
        way({lat1}, {lon1}, {lat2}, {lon2}) ["{tag}"];
        (._;>;);
        out body;
        """ 
    result = api.query(q) 

    lst = list()
    for way in result.ways: #TODO validate on IDE to interact with result if .ways is a valid attribute
        for node in way.nodes:
            lst.append((float(node.lat), float(node.lon))) 
    
    with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/' +  tag + r'_coordinates.csv', "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerows(lst)
    print(tag, 'extracted')

if False:
    extract_elements(lat1, lat2, lon1, lon2, 'highway')
    extract_elements(lat1, lat2, lon1, lon2, 'landuse') 
    extract_elements(lat1, lat2, lon1, lon2, 'natural')
    extract_elements(lat1, lat2, lon1, lon2, 'water')
    extract_elements(lat1, lat2, lon1, lon2, 'railway')
else:
    extract_elements(lat1, lat2, lon1, lon2, 'building')

# duds
# extract_elements(lat1, lat2, lon1, lon2, 'landcover=trees')
# extract_elements(lat1, lat2, lon1, lon2, 'wood')
# extract_elements(lat1, lat2, lon1, lon2, 'trees')
# extract_elements(lat1, lat2, lon1, lon2, 'landuse=forest')
# extract_elements(lat1, lat2, lon1, lon2, 'natural=wood')
import os
import argparse
import csv 

import overpy 

try: 
    parser = argparse.ArgumentParser() 
    parser.add_argument('-wdir', type=str, default=r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation', help="working dir")
    parser.add_argument('lat1', type=float, help="latitude1 for HGT file")
    parser.add_argument('lon1', type=float, help="longitude1 for HGT file")
    parser.add_argument('lat2', type=float, help="latitude2 for HGT file")
    parser.add_argument('lon2', type=float, help="longitude2 for HGT file")
    parser.add_argument('sample_size', type=int, help="squared root of samples")
    args = parser.parse_args() 

    os.chdir(args.wdir)  
except Exception as e:
    print(e)

lat1 = args.lat1
lon1 = args.lon1
lat2 = args.lat2
lon2 = args.lon2
sample_size = args.sample_size

p1 = (lat1, lon1)
p2 = (lat2, lon2)


def extract_elements(lat1, lat2, lon1, lon2, tag):
    api = overpy.Overpass()

    q = f"""
        way({lat1}, {lon1}, {lat2}, {lon2}) ["{tag}"];
        (._;>;);
        out body;
        """ 
    result = api.query(q) 

    lst = list()
    for way in result.ways:  
        
        for node in way.nodes:
            try: 
                lst.append((float(node.lat), float(node.lon), way.tags[tag])) 
            except:
                pass
    
    with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/' +  tag + r'_coordinates.csv', "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerows(lst)
    print(tag, 'extracted')

for tag in ['highway', 'landuse', 'natural', 'water', 'railway', 'building']:
    extract_elements(lat1, lat2, lon1, lon2, tag)
 
 
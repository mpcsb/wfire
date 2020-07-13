import os
import argparse
import csv 

import overpy 

 
lat1, lat2, lon1, lon2 = 38.718159, 38.78348, -9.481779 , -9.289
tag = 'highway'

api = overpy.Overpass()

q = f"""
    way({lat1}, {lon1}, {lat2}, {lon2}) ["{tag}"];
    (._;>;);
    out body;
    """ 
result = api.query(q) 

lst = []
for way in result.ways:  
    for node in way.nodes: 
        surface = ''
        if 'surface' in way.tags:
            surface = way.tags['surface'] 
            
        lst.append((float(node.lat), float(node.lon), way.tags[tag], surface)) 
# lst = list()
# for way in result.ways: #TODO validate on IDE to interact with result if .ways is a valid attribute
#     for node in way.nodes:
#         lst.append((float(node.lat), float(node.lon))) 

# with open(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/' +  tag + r'_coordinates.csv', "w", newline="") as f:
#     writer = csv.writer(f)
#     writer.writerows(lst)
# print(tag, 'extracted')

# if False:
#     extract_elements(lat1, lat2, lon1, lon2, 'highway')
#     extract_elements(lat1, lat2, lon1, lon2, 'landuse') 
#     extract_elements(lat1, lat2, lon1, lon2, 'natural')
#     extract_elements(lat1, lat2, lon1, lon2, 'water')
#     extract_elements(lat1, lat2, lon1, lon2, 'railway')
# else:
#     extract_elements(lat1, lat2, lon1, lon2, 'building')

# duds
# extract_elements(lat1, lat2, lon1, lon2, 'landcover=trees')
# extract_elements(lat1, lat2, lon1, lon2, 'wood')
# extract_elements(lat1, lat2, lon1, lon2, 'trees')
# extract_elements(lat1, lat2, lon1, lon2, 'landuse=forest')
# extract_elements(lat1, lat2, lon1, lon2, 'natural=wood')
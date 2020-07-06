import os
from math import floor
import gzip
import argparse
import csv 
# import numpy as np
from tqdm import tqdm
from gmalthgtparser import HgtParser   

try: 
    parser = argparse.ArgumentParser() 
    parser.add_argument('-wdir', type=str, default=r'C:\Users\Miguel\Documents\repos\wfire\elevation_data', help="working dir")
    parser.add_argument('lat1', type=float, help="latitude1 for HGT file")
    parser.add_argument('lon1', type=float, help="longitude1 for HGT file")
    parser.add_argument('lat2', type=float, help="latitude2 for HGT file")
    parser.add_argument('lon2', type=float, help="longitude2 for HGT file")
    # args = parser.parse_args()
    args, unknown = parser.parse_known_args()

    os.chdir(args.wdir) 
except Exception as e:
    print(e)
lat1 = args.lat1
lon1 = args.lon1
lat2 = args.lat2
lon2 = args.lon2

p1 = (lat1, lon1)
p2 = (lat2, lon2)


def linspace(lower, upper, num):  
    return [lower + x*(upper-lower)/num for x in range(num)]

def filename_gen(lat, lon):
    ''' receives coordinates used to select hgt file'''
    if lat >= 0: 
        fname = 'N'
        if lat >= 10:
            fname += str(lat)
        else:
            fname += '0' + str(lat)
    else:
        fname = 'S'
        if abs(lat) >= 10:
            fname += str(lat)
        else:
            fname += '0' + str(lat)
            
    if lon > 0:
        fname += 'E'
        if lon >= 100:
            fname +=str(lon)
        elif lon >= 10:
            fname += '0' + str(lon)
        else:
            fname += '00' + str(lon)
    else:
        fname += 'W'
        if abs(lon) >= 100:
            fname += str(abs(lon))
        elif abs(lon) >= 10:
            fname += '0' + str(abs(lon))
        else:
            fname += '00' + str(abs(lon))
    fname += '.hgt.gz'
    return  fname


def generate_topology(p1, p2, n_points=100):  
    np_lat = linspace(p1[0], p2[0], num=n_points )
    np_lon = linspace(p1[1], p2[1], num=n_points )
    
    coords_lat_lon = [(lat, lon) for lat in list(np_lat) for lon in list(np_lon)]
 
    lat = floor(p1[0])
    lon = floor(p1[1])
        
    filename = filename_gen(lat, lon)
    path = filename[:3] + '/' + filename
    decompressed_file = filename.replace('.gz', '')
     
    content = gzip.open(path).read()
        
    f = open(decompressed_file, 'wb')
    f.write(content)
    f.close()
    
     
    terrain = list()
    with HgtParser(decompressed_file) as parser:
        with tqdm(total=n_points**2) as pbar:
            for coord in coords_lat_lon:  
                lat_, lon_ = coord 
                alt = parser.get_elevation((lat_, lon_))
                terrain.append((lat_, lon_, alt[2]))
                pbar.update(1)
    os.remove(decompressed_file) 
    return terrain


terrain = generate_topology(p1, p2, n_points=80)

with open(r'C:\Users\Miguel\Documents\repos\wfire\wfire\src\simulation\terrain\temp\coords.csv', "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerows(terrain)

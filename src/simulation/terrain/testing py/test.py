# -*- coding: utf-8 -*-
"""
Created on Mon Jul 13 14:23:05 2020

@author: z003njns
"""

import pandas as pd
import os
import math
import matplotlib.pyplot as plt


os.chdir(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/')

lst = []
for f in os.listdir():
    lst.append([f,pd.read_csv(f)])
#%%

dfb = lst[0][1] # coords
dfb.columns=['lat', 'lon', 'type']
# dfb = dfb[dfb['type'] != 'yes']

df = lst[1][1] # coords
df.columns=['lat', 'lon', 'h']

dfr = lst[2][1] # roads
dfr.columns=['lat', 'lon', 'type']

dffff = lst[3][1] # landuse
dffff.columns=['lat', 'lon', 'type']
dffff = dffff[dffff['type'] == 'forest']


dfw = lst[6][1] # coords
dfw.columns=['lat', 'lon', 'type']

dff = lst[4][1] # roads
dff.columns=['lat', 'lon', 'type']
forest=['heath', 'scrub', 'wood','grassland', 'tree_row', 'spring']
dff = dff[dff['type'].isin(forest)]

plt.scatter(df['lon'], df['lat'], c=df['h'])
plt.scatter(dfr['lon'], dfr['lat'], c='black', s=1, alpha=0.01)
plt.scatter(dfw['lon'], dfw['lat'], c='blue', s=1)
plt.scatter(dfb['lon'], dfb['lat'], c='red', s=1, alpha=0.01)
plt.scatter(dff['lon'], dff['lat'], c='green', s=1, alpha=0.1)
plt.scatter(dffff['lon'], dffff['lat'], c='red', s=1, alpha=0.1)

 

df = lst[1][1] # coords
df.columns=['lat', 'lon', 'h']

df = df.sort_values(['lat', 'lon'], ascending = (True, True))
lats = list(df.lat)
lons = list(df.lon)
heights = list(df.h)

latitudes = sorted(list(set(lats)))
longitudes = sorted(list(set(lons)))

altitude_dict = {(v[0], v[1]): v[2] for v in  zip(lats, lons, heights)}


def adjacent_points(la, lo, h):
    i_lat = latitudes.index(la)
    i_lon = longitudes.index(lo)

    adjacent_coords = [(i,j) for i in [-1, 0, 1] for j in [-1, 0, 1]]
    adjacent = list()
    for i, j in adjacent_coords:
        if i == 0 and j == 0: continue
        try:
            latitude, longitude = latitudes[i_lat + i], longitudes[i_lon + j]
            adjacent.append((latitude, longitude, altitude_dict[latitude, longitude]))
        except:
            pass
    return adjacent

neighbours = list()
for coord in zip(lats, lons, heights):
    la, lo, h = coord
    n = adjacent_points(la, lo, h)
    neighbours.append([coord, n])


#%%

aspect = list()
for p1, n in neighbours:
    angles = list()
    for p2 in n:
        angles.append(angle(p1,p2))
    aspect.append([p1, angles])

for i in range(7):
    mean = list()
    for p, a in aspect:
        try: m = a[i]
        except : m = 0
        mean.append((p, m))


    x = [p[0][0] for p in mean]
    y = [p[0][1] for p in mean]
    z = [p[0][2] for p in mean]
    a = [p[1] for p in mean]

    plt.scatter(y, x, c=a, alpha=0.6, s=8)
    plt.show()
plt.scatter(y, x, c=z, alpha=0.7, s=8)

#%%
def dotproduct(v1, v2):
  return sum((a*b) for a, b in zip(v1, v2))

def length(v):
  return math.sqrt(dotproduct(v, v))

def angle(v1, v2):
  return math.degrees(math.acos(dotproduct(v1, v2) / (length(v1) * length(v2))))

p1, p2 = (38.709053999999995, -9.4812659375, 18),(38.804831625, -9.482466, 81)
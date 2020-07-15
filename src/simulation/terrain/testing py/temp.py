 

import os
os.chdir(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/testing py')

import richdem as rd
import numpy as np 


#%%

import pandas as pd

df = pd.read_csv(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/terrain/temp/coords.csv', header=None)
data = np.array(df)
n=600
m = rd.rdarray(np.flip(data[:,2].reshape(n,n), 0), no_data=-9999)



slope = rd.TerrainAttribute(m, attrib='slope_riserun')
rd.rdShow(slope, axes=False, cmap='jet', figsize=(8,5.5))


aspect = rd.TerrainAttribute(m, attrib='aspect')
rd.rdShow(aspect, axes=False, cmap='jet', figsize=(8,5.5))


profile_curvature = rd.TerrainAttribute(m, attrib='profile_curvature')
rd.rdShow(profile_curvature, axes=False, cmap='jet', figsize=(8,5.5))

curvature = rd.TerrainAttribute(m, attrib='curvature')
rd.rdShow(curvature, axes=False, cmap='jet', figsize=(8,5.5))

#%%
 

one_d = aspect.reshape(n**2)
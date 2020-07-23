import os 
import numpy as np  
import matplotlib.pyplot as plt   
import random
import pandas as pd
import imageio

os.chdir(r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/viz')


def getint(name): 
    num = name.split('.')[0] 
    return int(num)

files = [f for f in os.listdir( os.chdir)]
files_sorted = list(sorted(files, key=getint))


for f in files_sorted:
    df = pd.read_csv(f)
    plt.scatter(df.a, df.b, s=df.c, alpha=0.3)
    plt.savefig(str(i) + '.png')


images = [f for f in os.listdir(path) if 'png' in f]
images_sorted = list(sorted(files, key=getint))


images_lst = []
for filename in images_sorted:
    images_lst.append(imageio.imread(filename))
imageio.mimsave('animation.gif', images_lst)
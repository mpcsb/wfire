import os 
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D  
import numpy as np
import csv
import imageio

path = r'/home/miguel/Documents/projects/Wildfire/wfire/src/simulation/viz'
os.chdir(path)

try:
    os.remove('animation.gif')
except:
    pass

def getint(name): 
    ''' aux function used to sort integer strings'''
    num = name.split('.')[0] 
    return int(num)


def plot3d(f, i): 
    colors = {'road':'black',
              'water':'blue',
              'tree':'green',
              'burning_tree':'red',
              'ember':'orange',
              'ash':'grey'}

    content = list()
    with open(f) as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        line = 0
        for row in csv_reader:
            content.append(row[0].split(',')) 

    x_y_z = [(float(l[0]), float(l[1]), float(l[2])) for l in content[1:]]
    color = [colors[l[3]] for l in content[1:]]  
    x, y, z = list(map(list, zip(*x_y_z))) 


    x = np.array(x)#.reshape(self.terrain.num_points, self.terrain.num_points)
    y = np.array(y)#.reshape(self.terrain.num_points, self.terrain.num_points)
    z = np.array(z)#.reshape(self.terrain.num_points, self.terrain.num_points)
        
        
    ax = Axes3D(plt.figure(figsize=(15, 15)))
    ax.scatter(x, y, z, c=color, marker='^') 
    ax.view_init(35, 60 + i*10)  
    plt.savefig(str(i) + '.png')
    plt.close()



files = [f for f in os.listdir(path)]
files_sorted = list(sorted(files, key=getint)) 

for i, f in enumerate(files_sorted): 
    plot3d(f, i) 

images = [f for f in os.listdir(path) if 'png' in f]
images_sorted = list(sorted(images, key=getint))

images_lst = []
for filename in images_sorted:
    images_lst.append(imageio.imread(filename))
if len(images_lst) > 0:
    kargs = { 'duration': 0.15}
    imageio.mimsave('animation.gif', images_lst, **kargs) 

for filename in images_sorted + files_sorted:
    os.remove(filename)
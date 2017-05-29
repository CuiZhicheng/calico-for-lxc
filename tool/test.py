import subprocess,re,os,psutil,math,sys
import time,threading,json,traceback,platform
import env,os

def sys_run(command,check=False):
    Ret = subprocess.run(command, stdout = subprocess.PIPE, stderr = subprocess.STDOUT, shell=True, check=check)
    return Ret
 
def CreateContainer(container_name):
	#ret = subprocess.call("/home/pkusei/calico-for-lxc/lxc-create.sh %s" % container_name)
	ret = os.system("sh /home/pkusei/calico-for-lxc/lxc-create.sh %s" % container_name)
	print ret

def AttachContainerToNetwork(container_name, network_name):
	ret = os.system("sh /home/pkusei/calico-for-lxc/lxc-attach-calico.sh %s %s" % (container_name, network_name))
	print ret
	getIpByName(container_name)
	
def DetachContainerFromNetwork(container_name, network_name):
	ret = os.system("sh /home/pkusei/calico-for-lxc/lxc-detach-calico.sh %s %s" % (container_name, network_name))
	print ret

def DeleteContainer(container_name):
	ret = os.system("sh /home/pkusei/calico-for-lxc/lxc-delete.sh %s" % container_name)
	print ret

def getIpByName(container_name):
	output = subprocess.check_output("sudo lxc-info -n %s" % (container_name),shell=True)
        output = output.decode('utf-8')
        print output
        parts = re.split('\n',output)
        info = {}
        for part in parts:
            if not part == '':
                key_val = re.split(':',part)
                key = key_val[0]
                val = key_val[1]
                info[key] = val.lstrip()
        print info["IP"]	
	return info



if __name__ == '__main__':
	container_name = 't'
	network_name = 'frontend'
	CreateContainer(container_name)
	AttachContainerToNetwork(container_name, network_name)
	#container_info = getIpByName(container_name)
	#print container_info["IP"]

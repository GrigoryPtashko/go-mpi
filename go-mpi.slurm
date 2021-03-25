#!/bin/bash
#SBATCH --job-name=go-mpi        # create a short name for your job
#SBATCH --nodes=10                # node count
#SBATCH --ntasks=10               # total number of tasks across all nodes
#SBATCH --cpus-per-task=36        # cpu-cores per task (>1 if multi-threaded tasks)
##SBATCH --mem-per-cpu=1G         # memory per cpu-core (4G is default)
#SBATCH --time=00:01:00          # total run time limit (HH:MM:SS)
#SBATCH --mail-type=ALL
##SBATCH --mail-user=ptashko.go@talantiuspeh.ru
#SBATCH --mail-user=grigory.ptashko@gmail.com
#SBATCH --output="%x-%N-%j.out"
##SBATCH --test-only

module purge
module load gnu8/8.3.0
module load mpich/3.3

srun --mpi=pmi2 /home/ptashko/work/go-mpi/go-mpi

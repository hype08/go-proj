#!/usr/bin/python3

import os
import typer
from typing import *
from enum import Enum

DOCKER_OPTIONS = 'BUILDKIT=1'


class Environment(str, Enum):
  dev = 'dev'

root = os.getcwd()
app = typer.Typer()

@app.command(help='Builds Hype.')
def build():
  docker_execute(['build'])


@app.command(help='Starts Hype.')
def start():
  docker_execute(['up --remove-orphans'])


@app.command(help='Stops Hype.')
def stop():
  docker_execute(['down'])


@app.command(help='Open a shell into Hype.')
def shell(app: str):
  docker_execute(f'exec {app} bash')


@app.command(help='Reset the database.')
def reset_db():
  docker_execute_sql(['DROP DATABASE IF EXISTS hype WITH (FORCE)'])
  docker_execute_sql(['CREATE DATABASE hype'])


def docker_execute_sql(params: List[str]):
  '''Executes SQL in the database via a Docker Compose command.'''
  psql_params = ['host=db', 'user=postgres', 'password=postgres']
  command = f'run --rm --no-deps db psql "{" ".join(psql_params)}" -c "{" ".join(params)}"'
  docker_execute([command])


def docker_execute(command: List[str], env: Environment = Environment.dev):
  ''' Executes a Docker Compose command. '''
  cmd_str = ' '.join(command)
  os.system(f'{DOCKER_OPTIONS} ENV={env.value} docker compose -f {root}/docker-compose.yml {cmd_str}')


if __name__ == '__main__':
  app()

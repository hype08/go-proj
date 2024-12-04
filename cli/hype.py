#!/usr/bin/python3

import os
import typer

DOCKER_OPTIONS = 'BUILDKIT=1'

root = os.getcwd()
app = typer.Typer()

@app.command(help='Builds Hype.')
def build():
  docker_execute('build')


@app.command(help='Starts Hype.')
def start():
  docker_execute('up')


@app.command(help='Stops Hype.')
def stop():
  docker_execute('down')


@app.command(help='Open a shell into Hype.')
def shell(app: str):
  docker_execute(f'exec {app} bash')


def docker_execute(command: str):
  ''' Executes a Docker Compose command. '''
  os.system(f'{DOCKER_OPTIONS} docker-compose -f {root}/docker-compose.yml {command}')

if __name__ == '__main__':
  app()

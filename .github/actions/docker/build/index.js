const core = require('@actions/core');
const exec = require('@actions/exec');

async function run() {
    try {
        let name = core.getInput('name');
        let tag = core.getInput('tag');
        let fulltag = `${name}:${tag}`;
        const work_dir = core.getInput('working-directory')
        if (! work_dir) {
            work_dir = '.'
        }
        await exec.exec('docker build --tag',
        fulltag,
        '--file',
        core.getInput('dockerfile'),
        work_dir
        );
        isPushRequired = core.getInput('push');
        if (push) {
            await exec.exec('docker push',
            fulltag
            );  
        }
    }
    catch (err) {
        core.setFailed(err.message);
    }
}

run();


const core = require('@actions/core');
const exec = require('@actions/exec');

async function run() {
    try {
        await exec.exec('docker',[
            'login',
            '--username',
            core.getInput('username'),
            '--password',
            core.getInput('password')
        ]);
    }
    catch (err) {
        core.setFailed(err.message);
    }
}

run();


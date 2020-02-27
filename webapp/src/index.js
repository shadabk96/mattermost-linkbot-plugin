import Plugin from './plugin';
import manifest from './manifest';

window.registerPlugin(manifest.id, new Plugin());

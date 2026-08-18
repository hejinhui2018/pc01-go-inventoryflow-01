const fs = require('fs');
fs.mkdirSync('dist', { recursive: true });
const html = fs.readFileSync('index.html', 'utf8');
fs.writeFileSync('dist/index.html', html.replace('data-build="dev"', 'data-build="production"'));
console.log('built inventoryflow-console');

# om namah shivay

from flask import Flask
from flask import jsonify

app = Flask(__name__)

@app.route('/')
def hello():
    dictionary = {'message': 'Hello from Python!'}
    return jsonify(dictionary)

if __name__ == '__main__':
    app.run(debug=True)

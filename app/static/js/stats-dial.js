// stats-dial.js
class StatsDial extends HTMLElement {
    
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });

        // Default values and attributes
        this.value = parseInt(this.getAttribute('value')) || 0;
        this.min = parseInt(this.getAttribute('min')) || 0;
        this.max = parseInt(this.getAttribute('max')) || 100;

        // Create a canvas element within the shadow DOM
        const canvas = document.createElement('canvas');
        this.shadowRoot.appendChild(canvas);

        // Initialize the dial
        this.drawDial(canvas);

        // Watch for attribute changes
        const observer = new MutationObserver(mutations => {
            mutations.forEach(mutation => {
                if (mutation.type === 'attributes') {
                    this.value = parseInt(this.getAttribute('value')) || 0;
                    this.min = parseInt(this.getAttribute('min')) || 0;
                    this.max = parseInt(this.getAttribute('max')) || 100;
                    this.drawDial(canvas);
                }
            });
        });
        observer.observe(this, { attributes: true });
    }

    drawDial(canvas) {
        const ctx = canvas.getContext('2d');
        const width = this.clientWidth;
        const height = this.clientHeight;
        const centerX = width / 2;
        const centerY = height / 2;
        const radius = Math.min(centerX, centerY) * 0.8;

        // Clear canvas
        ctx.clearRect(0, 0, width, height);

        // Draw the dial background
        ctx.beginPath();
        ctx.arc(centerX, centerY, radius, 0, 2 * Math.PI);
        ctx.lineWidth = 10;
        ctx.strokeStyle = '#e0e0e0';
        ctx.stroke();

        // Draw the dial value
        const angle = (this.value - this.min) / (this.max - this.min) * 2 * Math.PI - 0.5 * Math.PI;
        ctx.beginPath();
        ctx.arc(centerX, centerY, radius, -0.5 * Math.PI, angle);
        ctx.lineWidth = 10;
        ctx.strokeStyle = '#4CAF50';
        ctx.stroke();
    }

    // Define attributes for value, min, and max
    static get observedAttributes() {
        return ['value', 'min', 'max'];
    }

    // Handle attribute changes
    attributeChangedCallback(name, oldValue, newValue) {
        if (name === 'value' || name === 'min' || name === 'max') {
            this.drawDial(this.shadowRoot.querySelector('canvas'));
        }
    }
}

// Define the custom element
customElements.define('stats-dial', StatsDial);

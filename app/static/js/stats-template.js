class StatsTemplate {

    constructor(data) {
        this.data = data;
    }
    
    generateStatsTemplate(statKey) {

        const stats = JSON.parse(this.data[statKey]);
        const statsJson = stats.stats_json;
        const title = this.formatTitle(statKey);

        
        let template = `<div class="${statKey}"><h2>${title}</h2>`;

        for (const [key, value] of Object.entries(statsJson)) {
            if (typeof value === 'object' && !Array.isArray(value) && value !== null) {
                template += this.generateNestedStatsTemplate(key,value);                
            } else {
                template += `<p>${this.formatTitle(key)}: ${value}</p>`;
            }
        }

        template += `</div>`;

        return template;

        return "ok";
        
    }

    generateNestedStatsTemplate(key, value) {
        let template = `<div class="${key}"><h3>${this.formatTitle(key)}</h3>`;
        for (const [nestedKey, nestedValue] of Object.entries(value)) {
            template += `<p>${this.formatTitle(nestedKey)}: ${nestedValue}</p>`;            
        }
        template += `</div>`;
        return template;
    }

    formatTitle(key) {
        return  key.replace(/_/g, ' ').replace(/\b\w/g, char => char.toUpperCase());
    }
    

    generateAllTemplates() {       
        let templates = '';
        for (const statKey of Object.keys(this.data)) {
            templates += this.generateStatsTemplate(statKey);
        }
        return templates;
    }
    
}

export default StatsTemplate;

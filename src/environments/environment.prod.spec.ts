import { TestBed } from '@angular/core/testing';
import { environment } from './environment';


describe('Testando environment PROD', () => {
    let url = environment.baseUrl;

    it('A URL tem que ser valida em baseurl', () => {
        expect(url).toBeDefined(); 
        expect(url).not.toBe('');  
        expect(() => new URL(url)).not.toThrow(); 
    });

});

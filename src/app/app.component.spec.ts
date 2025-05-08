import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppComponent } from './app.component';
import { NavigationEnd, Router } from '@angular/router';
import { of } from 'rxjs';
import { LoadingService } from './core/services/loading.service';
import { By } from '@angular/platform-browser';
import { HttpClientModule, provideHttpClient } from '@angular/common/http';
import { MatProgressSpinner } from '@angular/material/progress-spinner';
import { HttpService } from './core/services/http.service';
import { MatDialogRef } from '@angular/material/dialog';

describe('AppComponent', () => {
  let fixture: ComponentFixture<AppComponent>;
  let component: AppComponent;
  let mockRouter = { events: of(new NavigationEnd(0, '/', '/')) }; // Simula navegação
  let mockLoadingService = {
    isLoading$: of(false), // Simula um Observable
  };
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      providers: [
        provideHttpClient(), 
        { provide: Router, useValue: mockRouter },
        { provide: LoadingService, useValue: mockLoadingService },
        { provide: MatDialogRef, useValue: {} },
      ],
      imports: [AppComponent,HttpClientModule],
    }).compileComponents();
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
  });


  it("isLoading deve existir", ()=>{
    expect(fixture.componentInstance.isLoading).toBeDefined();
  });

  it("title deve existir", ()=>{
    expect(fixture.componentInstance.title).toBeDefined();
  });

  it("showHeader deve existir", ()=>{
    expect(fixture.componentInstance.showHeader).toBeDefined();
  });

  it('showHeader deve iniciar como falso', () => {
    expect(fixture.componentInstance.showHeader).toBeFalse();
  });

  it('isLoading deve iniciar como falso', () => {
    const fixture = TestBed.createComponent(AppComponent);
    expect(fixture.componentInstance.isLoading).toBeFalse();
  });

  it('deve criar o componente', () => {
    expect(component).toBeTruthy();
  });

  it('deve ter o Router e o LoadingService injetados', () => {
    expect(component['router']).toBeDefined();
    expect(component['loadingService']).toBeDefined();
  });

  it('showHeader nao deve ser exibido em login', () => {
    expect(component.showHeader).toBeFalse(); // Deve ser false se a URL for '/'
  });

  it('não deve renderizar <app-header> se showHeader for falso', () => {
    component.showHeader = false;
    fixture.detectChanges(); // Atualiza a view
    const headerElement = fixture.debugElement.query(By.css('app-header'));
    expect(headerElement).toBeNull();
  });

  it('deve renderizar <app-header> se showHeader for true', () => {
    component.showHeader = true;
    fixture.detectChanges(); // Atualiza a view
    const headerElement = fixture.debugElement.query(By.css('app-header'));
    expect(headerElement).toBeTruthy();
  });

  it('não deve renderizar <div class="loading-overlay"> se isLoading for falso', () => {
    component.isLoading = false;
    fixture.detectChanges();
    const loadingOverlay = fixture.debugElement.query(By.css('.loading-overlay'));
    expect(loadingOverlay).toBeNull();
  });

  it('deve renderizar <div class="loading-overlay"> se isLoading for true', () => {
    fixture.detectChanges();
    component.isLoading = true;
    fixture.detectChanges()
    const loadingOverlay = fixture.debugElement.query(By.css('.loading-overlay'));
    expect(loadingOverlay).toBeTruthy();
  });

  it('deve sempre renderizar <app-notificacao>', () => { 
    const notificacaoElement = fixture.debugElement.query(By.css('app-notificacao'));
    expect(notificacaoElement).toBeTruthy();
  });

  it('deve sempre renderizar <router-outlet>', () => {
    const routerOutlet = fixture.debugElement.query(By.css('router-outlet'));
    expect(routerOutlet).toBeTruthy();
  });
  
  
});

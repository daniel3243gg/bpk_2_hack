import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListaPedidoRecebidoComponent } from './lista-pedido-recebido.component';

describe('ListaPedidoRecebidoComponent', () => {
  let component: ListaPedidoRecebidoComponent;
  let fixture: ComponentFixture<ListaPedidoRecebidoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListaPedidoRecebidoComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListaPedidoRecebidoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

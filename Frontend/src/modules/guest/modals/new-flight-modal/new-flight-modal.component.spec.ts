import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewFlightModalComponent } from './new-flight-modal.component';

describe('NewFlightModalComponent', () => {
  let component: NewFlightModalComponent;
  let fixture: ComponentFixture<NewFlightModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewFlightModalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewFlightModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

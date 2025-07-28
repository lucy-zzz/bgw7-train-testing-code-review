package handler_test

import (
	"app/internal/handler"
	"app/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_FindByColorAndYearOK(t *testing.T) {
	// given
	sv := service.NewServiceVehicleDefaultStub()
	r := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("color", "red")
	rctx.URLParams.Add("year", "2010")
	expectedResult := 200
	w := httptest.NewRecorder()

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	// when

	h := handler.NewHandlerVehicle(sv).FindByColorAndYear()
	h.ServeHTTP(w, r)
	// then
	assert.Equal(t, expectedResult, w.Code)
}

func Test_FindByColorAndYearErrorParam(t *testing.T) {
	// given
	sv := service.NewServiceVehicleDefaultStub()
	r := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("color", "")
	rctx.URLParams.Add("year", "abc")
	expectedResult := http.StatusBadRequest
	w := httptest.NewRecorder()

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	// when

	h := handler.NewHandlerVehicle(sv).FindByColorAndYear()
	h.ServeHTTP(w, r)
	// then
	assert.Equal(t, expectedResult, w.Code)
}

func Test_FindByColorAndYearErrorService(t *testing.T) {
	// given
	sv := service.NewServiceVehicleDefaultStub()
	r := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("color", "blue")
	rctx.URLParams.Add("year", "2010")
	expectedResult := http.StatusInternalServerError
	w := httptest.NewRecorder()

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	// when

	h := handler.NewHandlerVehicle(sv).FindByColorAndYear()
	h.ServeHTTP(w, r)
	// then
	assert.Equal(t, expectedResult, w.Code)
}

func Test_FindByBrandAndYearRange(t *testing.T) {
	sv := service.NewServiceVehicleDefaultStub()

	t.Run("Should find by brand and year range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		rc := chi.NewRouteContext()
		rc.URLParams.Add("brand", "brand")
		rc.URLParams.Add("start_year", "2010")
		rc.URLParams.Add("end_year", "2011")
		expectedResult := http.StatusOK
		w := httptest.NewRecorder()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rc))

		//when
		h := handler.NewHandlerVehicle(sv).FindByBrandAndYearRange()
		h.ServeHTTP(w, r)

		// then
		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("Should throw error at start year param conversion when trying find by brand and year range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		rc := chi.NewRouteContext()
		rc.URLParams.Add("brand", "brand")
		rc.URLParams.Add("start_year", "")
		rc.URLParams.Add("end_year", "2011")
		expectedResult := http.StatusBadRequest
		w := httptest.NewRecorder()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rc))

		//when
		h := handler.NewHandlerVehicle(sv).FindByBrandAndYearRange()
		h.ServeHTTP(w, r)

		// then
		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("Should throw error at end year param conversion when trying find by brand and year range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		rc := chi.NewRouteContext()
		rc.URLParams.Add("brand", "brand")
		rc.URLParams.Add("start_year", "2010")
		rc.URLParams.Add("end_year", "")
		expectedResult := http.StatusBadRequest
		w := httptest.NewRecorder()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rc))

		//when
		h := handler.NewHandlerVehicle(sv).FindByBrandAndYearRange()
		h.ServeHTTP(w, r)

		// then
		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("Should throw error at service when trying find by brand and year range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		rc := chi.NewRouteContext()
		rc.URLParams.Add("brand", "err")
		rc.URLParams.Add("start_year", "2010")
		rc.URLParams.Add("end_year", "2011")
		expectedResult := http.StatusInternalServerError
		w := httptest.NewRecorder()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rc))

		//when
		h := handler.NewHandlerVehicle(sv).FindByBrandAndYearRange()
		h.ServeHTTP(w, r)

		// then
		assert.Equal(t, expectedResult, w.Code)
	})
}

func Test_AverageMaxSpeedByBrand(t *testing.T) {
	sv := service.NewServiceVehicleDefaultStub()
	t.Run("Should get average max speed", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "brand")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		expectedResult := http.StatusOK

		h := handler.NewHandlerVehicle(sv).AverageMaxSpeedByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("Should throw error at service when trying to get average max speed", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "err")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		expectedResult := http.StatusInternalServerError

		h := handler.NewHandlerVehicle(sv).AverageMaxSpeedByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("Should not find average max speed", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "not found")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		expectedResult := http.StatusNotFound

		h := handler.NewHandlerVehicle(sv).AverageMaxSpeedByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})
}

func Test_AverageCapacityByBrand(t *testing.T) {
	sv := service.NewServiceVehicleDefaultStub()
	t.Run("should get AverageCapacityByBrand", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "brand")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		expectedResult := http.StatusOK

		h := handler.NewHandlerVehicle(sv).AverageCapacityByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("should throw service internal error AverageCapacityByBrand", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "err")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		expectedResult := http.StatusInternalServerError

		h := handler.NewHandlerVehicle(sv).AverageCapacityByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})

	t.Run("should throw service not found error AverageCapacityByBrand", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "not found")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))

		expectedResult := http.StatusNotFound

		h := handler.NewHandlerVehicle(sv).AverageCapacityByBrand()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedResult, w.Code)
	})
}

func Test_SearchByWeightRange(t *testing.T) {
	sv := service.NewServiceVehicleDefaultStub()
	t.Run("should search by weight range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?weight_min=200&weight_max=250", nil)
		w := httptest.NewRecorder()
		expectedStatus := http.StatusOK

		h := handler.NewHandlerVehicle(sv).SearchByWeightRange()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("should throw weight_min convertion error when tries to search by weight range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?weight_min=asdas&weight_max=250", nil)
		w := httptest.NewRecorder()
		expectedStatus := http.StatusBadRequest

		h := handler.NewHandlerVehicle(sv).SearchByWeightRange()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("should throw weight_max convertion error when tries to search by weight range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?weight_min=1.15&weight_max=av", nil)
		w := httptest.NewRecorder()
		expectedStatus := http.StatusBadRequest

		h := handler.NewHandlerVehicle(sv).SearchByWeightRange()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("should throw error at service when tries to search by weight range", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?weight_min=0.0&weight_max=200", nil)
		w := httptest.NewRecorder()
		expectedStatus := http.StatusInternalServerError

		h := handler.NewHandlerVehicle(sv).SearchByWeightRange()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})
}
